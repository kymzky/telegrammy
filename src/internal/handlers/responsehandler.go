package handlers

import (
	"log/slog"
	"telegrammy/internal/config"
	"telegrammy/internal/domain"
	"telegrammy/internal/telegram"
	"telegrammy/internal/utils"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartPollingLoop() {
	pollInterval := time.Duration(config.GetPollInterval()) * time.Second
	lastUpdateID := 0
	updates := telegram.Poll(lastUpdateID)
	if len(updates) > 0 {
		lastUpdateID = updates[len(updates)-1].UpdateID
	}

	for {
		updates := telegram.Poll(lastUpdateID)
		for _, update := range updates {
			handleUpdate(update)
			lastUpdateID = update.UpdateID
		}
		time.Sleep(pollInterval)
	}
}

func handleUpdate(update tgbotapi.Update) {
	message := update.Message
	if message.Chat.ID != config.GetTelegramChatId() {
		return
	}
	responseMessage, parseMode := getResponseMessageAndParseMode(message)
	response := tgbotapi.NewMessage(config.GetTelegramChatId(), *responseMessage)
	response.ParseMode = *parseMode
	response.ReplyToMessageID = message.MessageID
	telegram.SendMessage(&response)
}

func getResponseMessageAndParseMode(message *tgbotapi.Message) (*string, *string) {
	content := message.Text
	slog.Info("Received new message.", "content", content)

	responseJobs := config.GetResponseJobs()
	for _, responseJob := range responseJobs {
		if response := getMessage(content, responseJob); response != nil {
			return response, &responseJob.ParseMode
		}
	}
	chatGptResponse := utils.GetChatGPTResponse(content)
	html := "HTML"
	return &chatGptResponse, &html
}

func getMessage(content string, responseJob domain.ResponseJob) *string {
	if responseJob.Trigger == content {
		return responseJob.GetMessage(config.GetShellPath())
	}
	return nil
}
