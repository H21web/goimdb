FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod first
COPY go.mod ./

# Download dependencies and create go.sum automatically
RUN go mod download && go mod tidy

# Copy entire source directory
COPY . ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o movie-api .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/movie-api .

EXPOSE 8080

CMD ["./movie-api"]
