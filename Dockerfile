# Base image with Go
FROM golang:1.22 as builder

# Set up working dir
WORKDIR /app

# Install build dependencies and clone source
RUN apt-get update && apt-get install -y git make curl

# Clone your node repo
RUN git clone https://github.com/ouroboros-crypto/new-node .

# Build and install the app
RUN make install

# --- Runtime image ---
FROM debian:bookworm-slim

# Install curl + basic tools
RUN apt-get update && apt-get install -y curl jq ca-certificates libc6 && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Create user
RUN useradd -m ouroboros

# Copy binary from builder
COPY --from=builder /go/bin/noded /usr/bin/ouroborosd

# Create necessary folders
RUN mkdir -p /home/ouroboros/.ouroboros && chown -R ouroboros:ouroboros /home/ouroboros

# Set user and working directory
USER ouroboros
WORKDIR /home/ouroboros
