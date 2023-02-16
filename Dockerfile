FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app

EXPOSE 9090

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]