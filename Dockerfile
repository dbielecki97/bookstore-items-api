FROM golang:alpine as builder

RUN apk update && apk upgrade && \
apk add --no-cache git
RUN apk add build-base
RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o items-api src/main.go

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/items-api .
EXPOSE 8083
CMD ["./items-api"]