FROM golang:1.19-alpine

WORKDIR /app

COPY . .
RUN apk add --update gcc musl-dev

RUN go mod download
RUN go build -o /stripeGo

EXPOSE 3000
CMD [ "/stripeGo" ]
