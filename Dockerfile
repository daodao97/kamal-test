FROM golang:1.22 AS builder

ARG VERSION

WORKDIR /build

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -X main.version=${VERSION}" -o app .

FROM alpine:latest AS final

WORKDIR /app
COPY *.yml /app/
COPY --from=builder /build/app /app/

RUN apk update && \
    apk add --no-cache sudo tzdata

# 设置所需的时区，例如亚洲/上海
ENV TZ=Asia/Shanghai

EXPOSE 8001

# 创建软链接，指向你想要的时区文件
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT ["/app/app"]