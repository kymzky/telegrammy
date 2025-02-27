FROM golang:1.24.0-alpine
# hadolint ignore=DL3018
RUN apk add --no-cache bash
WORKDIR /app
COPY . .
WORKDIR /app/cmd/telegrammy
RUN go build -o telegrammy
CMD ["./telegrammy"]
