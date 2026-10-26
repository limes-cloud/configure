# 构建阶段
# 注意：go.mod 中存在本地 replace 依赖（../manager/api、../../framework/kratosx），
# 构建前需在构建机上先执行 `go mod vendor` 将其固化到 vendor/ 目录，
# 构建时使用 -mod=vendor 避免访问上下文之外的本地路径。
FROM golang:alpine AS build
#ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
WORKDIR /src
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-s -w" -installsuffix=cgo -o /out/configure main.go

# 运行阶段
FROM alpine
WORKDIR /app
COPY --from=build /out/configure /app/configure
COPY ./conf/ /app/conf/
COPY ./deploy/ /app/deploy/
CMD ["./configure"]