FROM golang:1.17-alpine3.16 AS builder
LABEL stage=builder
WORKDIR /app

COPY . .
RUN go build -o main ./cmd/.

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY /web/ui /app/web/ui

EXPOSE 8080
CMD ["/app/main"]