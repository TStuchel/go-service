# The base go-image
FROM golang:latest

# Create a directory for the app
RUN mkdir /build

# Copy all files from the current directory to the app directory
COPY . /build

# Set working directory
WORKDIR /build

# Build
RUN go build ./cmd/go-service

# Run
ENTRYPOINT ["./go-service"]