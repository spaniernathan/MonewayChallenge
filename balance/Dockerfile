FROM golang:latest

COPY src/ "go/src/"
WORKDIR "go/src/"

RUN go get -d
RUN go build -o balance

EXPOSE 4040

CMD ["./balance"]