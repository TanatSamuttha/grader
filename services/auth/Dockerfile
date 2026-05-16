FROM golang:1.26 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

FROM debian:stable-slim
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 3000

CMD [ "./server" ]