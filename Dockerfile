# Stage 1: Get certificates and create user
FROM alpine AS certs
RUN apk add --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

# Create a minimal passwd file
RUN echo "app:x:1001:1001:App User:/home/app:/sbin/nologin" > /etc/passwd_minimal

# Stage 2: Build the Go application
FROM golang:latest AS builder
ARG GOARCH=arm64
WORKDIR /go/src/github.com/AntonioDaria/github-finder-golang 
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Use the build argument GOARCH to build for the specified architecture
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -mod=vendor -o app main.go

# Stage 3: Create the final minimal Docker image
FROM scratch
EXPOSE 8000
COPY --chown=1001:1001 --from=builder /go/src/github.com/AntonioDaria/github-finder-golang/app /githb-finder.service
COPY --chown=1001:1001 --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the minimal passwd file
COPY --from=certs /etc/passwd_minimal /etc/passwd

USER app
ENTRYPOINT ["/githb-finder.service"]
