FROM registry.new-page.xyz/golang:1.14.4-alpine

WORKDIR /go/src/app
COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk add --no-cache git \
    && go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["goproxy"]