FROM golang:1.18-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go


FROM alpine:latest AS production

WORKDIR /app
COPY --from=builder /app/main .
COPY  migrations ./migrations


CMD ["./main"]
