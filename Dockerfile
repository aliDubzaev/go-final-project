FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go build -o todo-server ./main.go

ENV TODO_PORT=7540
ENV TODO_DBFILE=/data/scheduler.db

VOLUME ["/data"]
EXPOSE 7540

CMD ["./todo-server"]