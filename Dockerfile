FROM golang:alpine AS builder
RUN apk add --no-cache make git
WORKDIR /app
COPY . .
EXPOSE 8080
CMD ["make"]
