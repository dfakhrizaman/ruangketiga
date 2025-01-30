# --------------------------------------------------
# 1) BUILD STAGE (for production builds)
# --------------------------------------------------
  FROM golang:1.23-alpine AS builder

  # Set working directory
  WORKDIR /app
  
  # Install necessary packages
  RUN apk update && apk add --no-cache git
  
  # Copy go.mod and go.sum first, then download modules
  COPY go.mod go.sum ./
  RUN go mod download
  
  # Copy the rest of your source code
  COPY . .
  
  # Build the binary
  RUN go build -o server ./cmd/server
  
  # --------------------------------------------------
  # 2) DEVELOPMENT STAGE (with Air hot reload)
  # --------------------------------------------------
    FROM golang:1.23-alpine as dev

    # Install the migrate CLI
    RUN apk update && apk add --no-cache git curl && \
        go install -v github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    
    RUN go install github.com/air-verse/air@latest

    ENV PATH="/go/bin:${PATH}"


    WORKDIR /app
    
    # Copy go.mod and go.sum first
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the entire project
    COPY . .
    
    # Copy the migrations folder (if you want it in the container)
    # This line is optional if you are mounting in docker-compose. 
    # But for a "pure Dockerfile" build, you need it.
    COPY migrations ./migrations
    
    # EXPOSE, etc.
    EXPOSE 8080
    
    CMD ["air"]
  
  # --------------------------------------------------
  # 3) PRODUCTION STAGE (final minimal image)
  # --------------------------------------------------
  FROM alpine:3.17 AS production
  
  WORKDIR /app
  
  # Copy the binary from the builder
  COPY --from=builder /app/server /app/server
  
  # Expose the port
  EXPOSE 8080
  
  # Run the binary
  CMD ["./server"]
  