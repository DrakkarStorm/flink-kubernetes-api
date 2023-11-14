# Build the manager binary
FROM golang:1.21 as builder
WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . .

# Build
RUN go build -a -o kubernetes-api .


# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
#FROM gcr.io/distroless/static:nonroot
FROM ubuntu:latest
WORKDIR /

COPY --from=builder /workspace/kubernetes-api .

RUN apt update && apt install -y curl

EXPOSE 9000
ENTRYPOINT ["/kubernetes-api"]
