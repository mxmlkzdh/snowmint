# Stage 1: Build the application using golang:alpine
FROM golang:1.23.1-alpine3.20 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o snowmint ./cmd/snowmint/main.go

# Stage 2: Create a minimal runtime image using scratch
FROM scratch

COPY --from=builder /app/snowmint /snowmint

ENTRYPOINT ["/snowmint"]
