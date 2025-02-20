FROM golang:1.24 as builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/app cmd/go-fiber-api/main.go

CMD ["/usr/local/bin/app","-config","config/local.yaml"]
