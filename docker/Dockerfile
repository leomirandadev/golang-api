###########DEBUG IMAGE###################
FROM golang:1.13-alpine3.10 as debug

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh \
        dpkg \
        gcc \
        git \
        musl-dev

WORKDIR /app
COPY . /app
RUN go test
RUN go build -o main

### Run the Delve debugger ###
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN go get github.com/go-delve/delve/cmd/dlv

COPY ./dlv.sh /
RUN chmod +x /dlv.sh 
ENTRYPOINT [ "/dlv.sh"]





###########PROD IMAGE###################
FROM alpine:3.9 as prod
COPY --from=debug /app/main /

CMD ./main