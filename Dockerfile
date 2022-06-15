# Build project's docker image
FROM golang:latest
RUN mkdir /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app

RUN go build -o main .
CMD ["/app/main"]