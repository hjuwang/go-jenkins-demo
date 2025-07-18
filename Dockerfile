FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go build -o app

FROM alpine:latest
COPY --from=builder /app/app /app
ENV VERSION=1.0.0
EXPOSE 8080
CMD ["/app"]
