FROM golang:1.15.2-alpine as dev

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apk add --no-cache alpine-sdk git \
    && go get github.com/pilu/fresh \
    && go get bitbucket.org/liamstask/goose/cmd/goose

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

CMD ["fresh"]


FROM golang:1.15.2-alpine as builder

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

# 本番用
FROM alpine as prod

WORKDIR /go/src/app

COPY --from=builder /go/src/app .

CMD ["./main"]