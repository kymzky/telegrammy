FROM golang:1.24.1-alpine
# hadolint ignore=DL3018
RUN apk add --no-cache bash
WORKDIR /app
COPY . .
WORKDIR /app/cmd/telegrammy
RUN go build -o telegrammy
CMD ["./telegrammy"]
