FROM golang:alpine AS build
ENV GO111MODULE=on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/build
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -installsuffix=cgo -o configure main.go

FROM alpine
WORKDIR /app
COPY --from=build /go/build/configure /app/configure
COPY ./conf/ /app/conf/
COPY ./deploy/ /app/deploy/
CMD ["./configure"]
