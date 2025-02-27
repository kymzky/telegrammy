FROM golang:1.24.0-alpine
RUN apk add --no-cache --upgrade \
  bash=5.2.37-r0
WORKDIR /app
COPY . .
WORKDIR /app/cmd/telegrammy
RUN go build -o telegrammy
CMD ["./telegrammy"]
