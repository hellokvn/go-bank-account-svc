FROM golang:1.18
WORKDIR /go/src/github.com/hellokvn/go-bank-account-svc
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]