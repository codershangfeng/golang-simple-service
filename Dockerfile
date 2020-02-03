FROM golang:alpine as build

# Uncomment later
# RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

WORKDIR /code

COPY ./src .

RUN go mod download

RUN go build -o main .

CMD ./main
