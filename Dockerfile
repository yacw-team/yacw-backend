# builder
FROM golang:1.20 as builder

# Copy project files into the builder
COPY . /server

# Set the working directory to the server directory
WORKDIR /server

# Install dependencies
RUN go get -d -v ./...

# Build the server
RUN make build

####################

# runner
FROM alpine:latest

# Create a directory for the server
RUN mkdir /go

# Set the working directory to the server directory
WORKDIR /go

# Copy the executable server and database file into the container
COPY --from=builder /server/server /go

# Expose the port that the server will be running on
EXPOSE 8080

# Run the server
CMD ["./server"]
