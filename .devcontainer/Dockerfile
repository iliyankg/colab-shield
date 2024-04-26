ARG VARIANT="0-1.22-bookworm"

FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT}

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive && apt-get install -y protobuf-compiler

USER vscode

# Install Go Protobuf tools.
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/spf13/cobra-cli@latest

# Install Docker tools 
# For Docker in Docker dev work.
RUN curl -sSL https://get.docker.com/ | sh
