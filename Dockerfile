FROM golang:latest

#RUN go get .
ADD golang-chat-server /go/golang-chat-server
ENTRYPOINT /go/golang-chat-server

# Service listens on port 8084.
EXPOSE 8084
