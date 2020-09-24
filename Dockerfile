FROM golang:alpine AS builder
RUN echo $GOPATH
RUN apk add build-base --no-cache --update-cache --repository https://mirrors.aliyun.com/alpine/latest-stable/main/ --allow-untrusted
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"
RUN adduser -D -g 'root' appuser
WORKDIR /go/src/gitlab.xpaas.lenovo.com/hcmp/overseer
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o /bin/overseer

FROM alpine
RUN apk add tzdata --no-cache --update-cache --repository https://mirrors.aliyun.com/alpine/latest-stable/main/ --allow-untrusted && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
WORKDIR /data/overseer
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /bin/overseer .
RUN chown -R appuser:root /data/overseer
USER appuser
ENTRYPOINT [ "./overseer" ]