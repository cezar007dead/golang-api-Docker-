FROM golang:latest

LABEL maintainer="Quique <gurgen.meliksetyants@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

COPY base.json .

RUN go mod download

COPY . .

RUN go build

CMD ["./server"]