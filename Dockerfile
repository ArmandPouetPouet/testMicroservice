FROM golang:latest

ADD . /go/src/testMicroservice
WORKDIR /go/src/testMicroservice
RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go

RUN go install testMicroservice

ENTRYPOINT /go/bin/testMicroservice

EXPOSE 8080

#docker run -d -p 127.0.0.1:8080:8080 testmicroservice 