FROM golang:alpine AS builder
RUN apk add --no-cache make git
WORKDIR /app
COPY . .
EXPOSE $PORT
CMD ["make"]
