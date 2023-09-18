FROM golang:1.21.0

WORKDIR /go/cmd/

COPY . .

RUN go build -o app ./cmd/

EXPOSE 3033 

CMD ["./app"]
