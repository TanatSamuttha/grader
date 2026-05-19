FROM golang:1.26 AS builder
WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY ./src .

RUN go build -o server .

FROM debian:stable-slim
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 3001

CMD [ "./server" ]