FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

WORKDIR /build/server

# Build the application
RUN go build -o main .

# Test the application
RUN go test ./...

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/server/main .

# Build a small image
FROM alpine

# Update certificates - certain services won't relay information properly if these aren't set
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates

COPY --from=builder /dist/main /

# Command to run
ENTRYPOINT ["/main"]