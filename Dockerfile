# Build the manager binary
FROM golang:1.21 as builder

# Définis le répertoire de travail
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . .

# Construis l'application
RUN CGO_ENABLED=0 GOOS=linux go build -o flink-kubernetes-api

# Utilise une image légère d'Alpine pour exécuter l'application
FROM alpine:latest

# Définis le répertoire de travail
WORKDIR /app

# Copie l'exécutable construit à partir de l'étape précédente
COPY --from=builder /app/flink-kubernetes-api .

# Expose le port sur lequel l'application écoute
EXPOSE 8080

# Exécute l'application
CMD ["./flink-kubernetes-api"]
