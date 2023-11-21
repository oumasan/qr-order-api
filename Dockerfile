FROM golang:latest

WORKDIR /app

#RUN go mod tidy

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]

EXPOSE 9090