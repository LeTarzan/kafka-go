FROM golang:1.22 AS builder

WORKDIR /app

COPY ./consumer .
RUN go mod download
RUN CGO_ENABLED=0 go build -o main .
RUN chmod +x ./main

FROM debian:buster-slim

WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]
