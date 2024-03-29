# First stage: build the executable.
# Start from the official Go image to create a build artifact.
# This is based on Debian and includes standard C libraries.
FROM golang:1.21 AS builder

ARG PROTOC_VERSION=25.2

RUN apt-get update && apt-get install -y unzip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /usr/src/app

# Install protoc
RUN wget -nv https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-aarch_64.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-aarch_64.zip -d /usr/local && \
    rm protoc-${PROTOC_VERSION}-linux-aarch_64.zip
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH $PATH:$(go env GOPATH)/bin

# Copy the Go Modules manifests and download the dependencies.
# This is done before copying the code to leverage Docker cache layers.
COPY go.* ./
RUN go mod download

# Copy the source code from the current directory to the working directory inside the container.
COPY . .

# Create a non-root user and group to run the application.
RUN groupadd -r nonroot && useradd --no-log-init -r -g nonroot nonroot

# Build the binary with full module support and without Cgo.
# Compile the binary statically including all dependencies.
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -a -installsuffix cgo -o /go/bin/main ./cmd/observer

# Second stage: build the runtime container.
# Start from a scratch image, which is an empty container.
FROM scratch AS runtime

WORKDIR /usr/src/app

# Import the the user and group information
COPY --from=builder /etc/passwd /etc/group /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
# This is important for applications that make external HTTPS calls.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled binary from the first stage.
COPY --from=builder /go/bin/main .

# Use the nonroot user to run the application
USER nonroot:nonroot

# Declare the environment variable for the application.
# For example, setting the port where the application will run.
# ENV PORT=8080

# Expose the application on port 8080.
# EXPOSE $PORT

# Define the entry point for the docker image.
# This is the command that will be run when the container starts.
ENTRYPOINT ["/usr/src/app/main"]
