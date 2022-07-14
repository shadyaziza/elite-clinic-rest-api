FROM golang:1.18-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# if migrate is used as CLI tool
#RUN apk add curl
#RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app/app .
# if migrate is used as CLI tool
#COPY --from=builder /app/migrate .
COPY  start.sh .
COPY  migrations ./migrations
COPY wait-for.sh .

# this is overriden by docker compose command arguemnts
CMD ["app/app"]
# combining CMD with ENTRYPOINT is equivalent to
# ENTRYPOINT ["/start.sh","./app"]
ENTRYPOINT ["/app/start.sh"]