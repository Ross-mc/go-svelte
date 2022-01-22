# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

EXPOSE 443

EXPOSE 8080

HEALTHCHECK CMD curl -f http://localhost:443/health || exit 1

# Command to run the executable
CMD ["./main"]
#