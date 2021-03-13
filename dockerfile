# 编译镜像
FROM golang:1.15.3 AS builderImage
WORKDIR /go/src/project
COPY . .
# 配置免密
RUN echo 'https://username:password@dev.phoenix-ea.com' > ~/.git-credentials
RUN git config --global credential.helper store
RUN go env -w GO111MODULE=on
RUN go env -w GONOPROXY=dev.phoenix-ea.com
RUN go env -w GOPRIVATE=dev.phoenix-ea.com
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go mod init
RUN go mod tidy
RUN go build -o application  .
# 运行镜像
FROM alpine:3.12.0
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezon && rm -rf /tmp/* \
    && rm -rf /var/cache/apk/* && mkdir -p /root/project
WORKDIR /root/project 
COPY --from=builderImage /go/src/project/application .
COPY --from=builderImage /go/src/project/etc .
EXPOSE 8080 
ENTRYPOINT [ "./application" ]