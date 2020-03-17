FROM golang:1.12-alpine

# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Leonardo <leonardo@planoviver.com>"

# dependencie crypto
RUN go get golang.org/x/crypto/bcrypt

# dependencie http
RUN go get -u github.com/gorilla/mux

# dependencie ORM
RUN go get -u github.com/jinzhu/gorm

# dependencie MYSQL
RUN go get -u github.com/go-sql-driver/mysql

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]