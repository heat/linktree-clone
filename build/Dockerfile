FROM golang:latest

LABEL maintainer="Onezino Moreira <zinogabriel@gmail.com>"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main cmd/linktree/main.go

EXPOSE 8080

CMD ["./main"]
