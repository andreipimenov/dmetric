FROM golang:1.10

RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/github.com/andreipimenov/dmetric

WORKDIR /go/src/github.com/andreipimenov/dmetric

RUN dep ensure

RUN go build -o /go/bin/monitor cmd/monitor/*.go

ENTRYPOINT ["/go/src/github.com/andreipimenov/dmetric/entrypoint.sh"]

CMD monitor