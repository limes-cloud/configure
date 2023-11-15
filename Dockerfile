FROM golang:alpine AS build

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/build
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o configure cmd/configure/main.go

FROM alpine
EXPOSE 7080
EXPOSE 7081
EXPOSE 7082
WORKDIR /go/build
COPY ./config/config-prod.yaml /go/build/config/config.yaml
COPY ./static /go/build/static
COPY ./web/dist /go/build/web/dist
COPY --from=build /go/build/configure /go/build/configure
CMD ["./configure"]
