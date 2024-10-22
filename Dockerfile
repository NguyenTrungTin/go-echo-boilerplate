FROM golang:1.13.5-alpine

RUN apk update && apk upgrade && \
	apk add --no-cache bash git openssh

WORKDIR /app

RUN go get github.com/pilu/fresh

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 80

CMD ["./main"]
