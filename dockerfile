FROM golang:1.19 AS builder
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
