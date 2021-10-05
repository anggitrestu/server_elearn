# stage I - khusus build dengan envinroment yang sama
FROM golang:1.16-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main


FROM alpine:3.14
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .
EXPOSE 8080
CMD ["./main"]
