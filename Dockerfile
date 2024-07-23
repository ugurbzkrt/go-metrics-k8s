FROM golang:1.22-alpine

WORKDIR /basics

COPY . .

RUN go build -o main .

CMD ["./main"]
