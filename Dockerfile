FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o receipt-processor cmd/main.go

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/receipt-processor .

RUN chown appuser:appgroup /app/receipt-processor

USER appuser

EXPOSE 8080

CMD ["./receipt-processor"]
