FROM golang:1.17.3 as development

ENV MYSQL_USER=user \
    MYSQL_PASSWORD=password \
    MYSQL_PORT=3306 \
    MYSQL_DB=challengedb \
    MYSQL_SERVICE_NAME=db \
    FIREBASE_API_KEY=

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

# Copy app files
COPY . .

# Expose port
EXPOSE 4000

# Run the executable
CMD go run ./... --start-service
