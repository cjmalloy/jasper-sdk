FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -a -installsuffix cgo -o /app/main .

FROM alpine:latest as deploy
WORKDIR /app
COPY --from=builder /app/main .
RUN chmod +x main
ENTRYPOINT /app/main
