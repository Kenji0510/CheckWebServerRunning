FROM golang:1.22-alpine

RUN apk add --no-cache bash git

WORKDIR /app

RUN git clone https://github.com/Kenji0510/CheckWebServerRunning.git .

RUN go mod download

RUN go build -o main .

COPY .env .env

CMD ["./main"]