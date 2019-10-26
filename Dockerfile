FROM golang:latest
WORKDIR /app
COPY . .

RUN go build -o /ontheroad
EXPOSE 9000

CMD ["/ontheroad", "migrate","start"]