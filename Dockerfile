# Start from an official Go runtime base image
FROM golang:1.22

# Set the current working directory in the container
WORKDIR /app

# Copy Go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working directory in the container
COPY . .

# Build the application
RUN go build -o main .

EXPOSE 8080

# Command to run the application
CMD ["./main"]
