FROM golang as builder

COPY src/ /go/build/
WORKDIR /go/build/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o init .
