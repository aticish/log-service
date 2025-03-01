FROM golang:1.24-alpine
RUN apk update && apk add --no-cache ca-certificates

COPY ./api .

RUN go mod download

EXPOSE 3000

CMD ["go", "run", "cmd/main.go"]
