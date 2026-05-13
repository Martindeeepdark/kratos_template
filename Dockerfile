FROM golang:1.24-alpine AS builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o kratos-template ./cmd


FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/kratos-template .
COPY --from=builder /app/configs ./configs

EXPOSE 8000 9000

ENTRYPOINT ["./kratos-template"]
CMD ["greeter-service", "--config", "configs/config.yaml"]
