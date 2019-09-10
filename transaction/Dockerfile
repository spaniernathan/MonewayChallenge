FROM golang:latest

COPY src/ "go/src/"
WORKDIR "go/src/"

RUN go get -d
RUN go build -o transaction

EXPOSE 8080

CMD ["./transaction"]