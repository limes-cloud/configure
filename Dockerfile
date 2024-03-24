# 构建go
FROM golang:alpine AS gobuild
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/build
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o configure cmd/configure/main.go

#构建web
FROM gplane/pnpm:8.9.2-node18 AS webbuild
WORKDIR /app/
ADD web/package.json /app/
RUN pnpm config set registry=https://registry.npmmirror.com/
RUN pnpm install
ADD ./web/ /app/
RUN pnpm build

FROM alpine
WORKDIR /go/build
COPY ./internal/config/config-prod.yaml /go/build/internal//config/config.yaml
COPY ./static /go/build/static
COPY --from=webbuild /app/dist/ /go/build/web/dist/
COPY --from=gobuild /go/build/configure /go/build/configure
CMD ["./configure"]
