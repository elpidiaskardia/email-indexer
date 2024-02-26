#Golang image

FROM golang:1.22.0-alpine

#Env vars
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GO_ENV=production

## Project setup

# Create app directory inside container
RUN mkdir app

# cd /app
WORKDIR /app

COPY go.mod .

# Download all dependencies
RUN go mod download

# Declare custom args
ARG DB_USER \
    DB_PASSWORD \
    DB_PORT

ENV DB_USER=${DB_USER} \
    DB_PASSWORD=${DB_PASSWORD} \
    DB_PORT=${DB_PORT}

# Copies all files from local to container
COPY . .

# Build
RUN go build -o main .

# Run
CMD ["/app/main"]