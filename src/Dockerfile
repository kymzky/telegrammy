FROM golang:1.24.0-alpine
# renovate: datasource=repology depName=alpine/bash versioning=loose
RUN apk add --no-cache \
  bash=5.2.37-r0
WORKDIR /app
COPY . .
WORKDIR /app/cmd/telegrammy
RUN go build -o telegrammy
CMD ["./telegrammy"]
