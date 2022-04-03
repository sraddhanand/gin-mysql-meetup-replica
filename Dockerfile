FROM golang:alpine as builder
WORKDIR /app
ENV GO111MODULE=on
COPY ./gin /app/gin
RUN apk add --no-cache git && \
    cd /app/gin && \
    go mod tidy && \
    go build -o ./meetup .

FROM alpine:latest
RUN adduser -D -s /bin/sh -h /home/apprunner apprunner
USER apprunner
WORKDIR /home/apprunner
COPY --from=builder /app/gin/meetup .
EXPOSE 8080
CMD ["./meetup"]