# Specifies a parent image
FROM golang:1.22.3-bookworm

# Creates an app directory to hold your app's source code
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY ./app/go.mod ./

# Install go dependencies
RUN go mod download

# Copy the rest of the application code and the CSV file
COPY ./app ./
COPY ./Crime_Data_from_2020_to_Present.csv .

# Build with your optional configuration
RUN go build -o /crime-stats ./cmd/main.go

# Tell docker which network port to expose 
EXPOSE 8080

# Specify the executable command that runs when the container starts
CMD [ "/crime-stats" ]