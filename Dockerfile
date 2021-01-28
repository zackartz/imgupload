FROM golang:1.15 AS build

WORKDIR /go/src/imgupload

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/imgupload/app .

EXPOSE 1337

CMD ["./app", "-prod"]
