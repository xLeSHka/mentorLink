FROM golang:1.23-alpine3.21 AS builder

# Setup base software for building an app.
RUN apk update && apk add ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt

# Download dependencies.
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy application source.
COPY . .

# Build the application.
RUN go build -o bin/application ./cmd/main.go

# Prepare executor image.
FROM alpine:3.21 AS runner

RUN apk update && apk add ca-certificates libc6-compat openssh bash && rm -rf /var/cache/apk/*

WORKDIR /opt

COPY migrations migrations

COPY --from=builder /opt/bin/application ./
COPY ./.env ./
# Run the application.
CMD ["./application"]