FROM golang:1.21.0

WORKDIR /go/cmd/app/

COPY . .

RUN go build -o app ./cmd/app

EXPOSE 3033 

CMD ["./app"]
