package telegram

import (
	"log/slog"
	"sync"
	"telegrammy/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	telegramBot *tgbotapi.BotAPI
	once        sync.Once
)

func Poll(updateId int) []tgbotapi.Update {
	telegramBot := getTelegramBot()
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Offset = int(updateId) + 1
	updates, err := telegramBot.GetUpdates(updateConfig)
	if err != nil {
		slog.Error("Failed to poll messages.", "err", err)
	}
	return updates
}

func SendMessage(message *tgbotapi.MessageConfig) {
	telegramBot := getTelegramBot()
	_, err := telegramBot.Send(message)
	if err != nil {
		slog.Error("Failed to send telegram message.", "err", err, "message", message)
	}
}

func getTelegramBot() *tgbotapi.BotAPI {
	once.Do(func() {
		var err error
		telegramBotToken := config.GetTelegramBotToken()
		telegramBot, err = tgbotapi.NewBotAPI(telegramBotToken)
		if err != nil {
			slog.Error("Failed to initialize telegram bot.", "err", err)
			return
		}
	})
	return telegramBot
}
