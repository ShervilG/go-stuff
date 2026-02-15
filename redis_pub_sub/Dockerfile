FROM golang:1.23-alpine

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -o main .
RUN apk add --no-cache redis

CMD sh -c "redis-server & ./main"

