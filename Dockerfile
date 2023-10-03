FROM golang:1.21

WORKDIR /usr/src/app
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app cmd/main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
CMD ["app"]