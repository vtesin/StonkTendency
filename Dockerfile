FROM golang:1.16.0-alpine

ENV LIBRARY_ENV=dev
ENV CGO_ENABLED=0
ENV GOOS=linux

ADD . /opt/go/src/local/stonktendency
WORKDIR /opt/go/src/local/stonktendency

RUN go get github.com/go-delve/delve/cmd/dlv
RUN go build -a -installsuffix cgo -tags "$LIBRARY_ENV netgo" -installsuffix netgo -o api/main api/main.go
CMD ["api/main"]