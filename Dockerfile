FROM golang:1.12

WORKDIR /go/src/GOREsT
COPY src/* ./

RUN go get -v

CMD GOREsT
