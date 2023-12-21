FROM golang:alpine as builder

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o main main.go

FROM alpine

COPY --from=builder /usr/src/app/main /app

EXPOSE 8000

ENTRYPOINT [ "/app" ]