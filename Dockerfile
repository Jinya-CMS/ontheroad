FROM golang:1.13-alpine
WORKDIR /app
COPY . .

RUN go build -o /ontheroad
EXPOSE 9000

CMD ["/ontheroad", "migrate", "start"]