FROM alpine:3.21.2
RUN apk add --no-cache \
  bash=5.2.37-r0 \
  go=1.23.6-r0
WORKDIR /app
COPY . .
WORKDIR /app/cmd/telegrammy
RUN go build -o telegrammy
CMD ["./telegrammy"]
