FROM golang:1.22

WORKDIR /usr/src/api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/api ./...

CMD ["api"]
