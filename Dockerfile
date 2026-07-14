# ==========================
# Stage 1 - Build
# ==========================
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o be-dashboard .

# ==========================
# Stage 2 - Runtime
# ==========================
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/be-dashboard .
COPY --from=builder /app/.env .
COPY --from=builder /app/firebase-service-account.json .

EXPOSE 3000

CMD ["./be-dashboard"]
