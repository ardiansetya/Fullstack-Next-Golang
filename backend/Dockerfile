# Stage 1: Builder
# Menggunakan versi Go terbaru yang stabil
FROM golang:1.22.6-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Salin go.mod dan go.sum
COPY go.mod go.sum* ./

# Set GOTOOLCHAIN untuk mengatasi masalah versi
ENV GOTOOLCHAIN=local

# Download dependencies
RUN go mod download

# Salin seluruh source code
COPY . .

# Build aplikasi
RUN go build -o /app/main .

# Stage 2: Final Image
FROM alpine:3

WORKDIR /app

# Salin binary dari builder stage
COPY --from=builder /app/main .

CMD ["/app/main"]