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

CMD ["/app/dist/main"]
