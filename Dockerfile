FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o realtime-analytics ./cmd/server

# Final small image
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/realtime-analytics .

EXPOSE 8080

CMD ["/app/realtime-analytics"]