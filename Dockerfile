FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . .
RUN go mod tidy
RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/app /usr/src/app/cmd/cats_api


CMD ["app"]
