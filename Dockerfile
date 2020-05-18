ARG GO_VERSION=1.14

FROM golang:$GO_VERSION as builder

MAINTAINER zpp

# 交叉编译不支持 cgo ,禁用CGO_ENABLED
ENV GOPROXY=https://goproxy.io \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o seckill

#ENTRYPOINT ["./seckill"]
FROM scratch

WORKDIR /app/seckill
COPY --from=builder /build ./
EXPOSE 8000
ENTRYPOINT ["./seckill"]