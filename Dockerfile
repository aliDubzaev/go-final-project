FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go build -o todo-server ./main.go

VOLUME ["/data"]

EXPOSE 7540

CMD ["./todo-server"]