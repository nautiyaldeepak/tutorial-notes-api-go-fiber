FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o goapp .

FROM alpine:3.19.1
WORKDIR /root/
COPY --from=builder /app/goapp .
EXPOSE 3000
EXPOSE 2112
EXPOSE 8080
CMD ["./goapp"]