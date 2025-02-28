package config

import (
	"log/slog"
	"os"
	"strconv"
)

const (
	telegramBotToken = "TELEGRAM_BOT_TOKEN"
	telegramChatId   = "TELEGRAM_CHAT_ID"
	openAiApiKey     = "OPENAI_API_KEY"
)

func GetTelegramBotToken() string {
	return getStringValue(telegramBotToken)
}

func GetTelegramChatId() int64 {
	return getInt64Value(telegramChatId)
}

func GetOpenAiApiKey() string {
	return getStringValue(openAiApiKey)
}

func getStringValue(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		slog.Error("Environment variable not set.", "key", key)
	}
	return value
}

func getInt64Value(key string) int64 {
	stringValue := getStringValue(key)
	value, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		slog.Error("Failed to parse int64.", "stringValue", stringValue)
	}
	return value
}
