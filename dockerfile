FROM golang:1.23.5

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o GoMux .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./GoMux"]