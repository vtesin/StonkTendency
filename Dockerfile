FROM golang:1.16.0-alpine
WORKDIR /bin
COPY ./bin .
CMD ["api"]