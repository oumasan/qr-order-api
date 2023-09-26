FROM golang:latest

WORKDIR /src
COPY ./src /src

RUN go mod init main \
  && go mod tidy \
  && go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 9090


CMD ["go", "run", "main.go"]