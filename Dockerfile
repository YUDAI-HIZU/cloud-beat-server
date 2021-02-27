FROM golang:1.15.2-alpine as dev

WORKDIR /app

ENV GO111MODULE=on

RUN apk add --no-cache alpine-sdk git \
    && go get github.com/pilu/fresh \
    && go get bitbucket.org/liamstask/goose/cmd/goose

COPY . .

CMD ["fresh"]

FROM golang:1.15.2-alpine as builder

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .


FROM alpine as prod

WORKDIR /app

COPY --from=builder /src/main .

CMD ["./main"]