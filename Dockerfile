FROM golang:1.16 as builder

# Create and change to the app directory.
WORKDIR /app

RUN apt-get update # Last Modified: 2021-04-19T21:28:04
RUN apt-get install -y xorg-dev
RUN apt-get install -y libgl1-mesa-dev

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go install -mod=readonly -v ./cmd/golozd

# Run the web service on container startup.
CMD ["/go/bin/golozd"]

