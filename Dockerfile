# The base go-image
FROM golang:alpine

# Copy all of the build files into a /build directory
WORKDIR /build
COPY ./ ./

# Build
RUN go build ./cmd/go-service

# Copy executable / cleanup build
WORKDIR ../app
RUN mv ../build/go-service go-service \
    && rm -R ../build \
    && go clean -cache -modcache -i -r

# Run
ENTRYPOINT ["./go-service"]