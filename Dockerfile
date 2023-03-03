FROM golang:1.18.1-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o test-assignment-gocloudcamp ./cmd/server/main.go

CMD [ "./test-assignment-gocloudcamp" ]