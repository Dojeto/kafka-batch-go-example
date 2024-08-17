FROM golang:1.22

WORKDIR /app

COPY go.mod go.mod

COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build producer/*.go

RUN go build consumer/*.go

CMD ./kafka & ./db