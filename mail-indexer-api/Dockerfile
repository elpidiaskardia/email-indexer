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
ARG ZINC_URL \
    AUTH_USERNAME \
    AUTH_PASSWORD \
    DB_PORT

ENV ZINC_URL=${ZINC_URL} \
    AUTH_USERNAME=${AUTH_USERNAME} \
    AUTH_PASSWORD=${AUTH_PASSWORD} \
    DB_PORT=${DB_PORT}

# Copies all files from local to container
COPY . .
# Run the main.go in the db folder to load the database

# Build
RUN go build -o main .

# Run
CMD ["/app/main"]