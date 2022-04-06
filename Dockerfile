FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./bin/app ./cmd/main.go

FROM alpine

COPY --from=build ./app/bin/app ./bin/app

CMD ["./bin/app"]