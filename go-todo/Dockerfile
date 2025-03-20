# Build stage
FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/dist/main .

# Runtime stage
FROM golang:1.24

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/dist/main /app/dist/main

# Copy the source code (optional)
COPY . .

# Install SQLite (needed for debugging)
RUN apt-get update && apt-get install -y sqlite3

CMD ["/app/dist/main"]
