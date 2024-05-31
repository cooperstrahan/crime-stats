# Specifies a parent image
FROM golang:1.22.3-bookworm

# Creates an app directory to hold your app's source code
WORKDIR /app

COPY ./app .

# Install go dependencies
RUN go mod download

# Build with your optional configuration
RUN go build -o /crime-stats ./cmd/main.go

# Tell docker which network port to expose 
EXPOSE 8080

# Specify the executable command that runs when the container starts
CMD [ "/crime-stats" ]