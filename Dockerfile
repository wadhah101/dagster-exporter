FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o exporter .

FROM alpine:3.20.3

WORKDIR /root/

COPY --from=builder /app/exporter .

EXPOSE 8080

ENTRYPOINT ["./exporter"]
