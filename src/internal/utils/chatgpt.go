package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"telegrammy/internal/config"

	"github.com/sashabaranov/go-openai"
)

const (
	assistant = "Assistant"
)

func GetChatGPTResponse(message string) string {
	conversation := getConversation()
	userMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	}
	conversation = append(conversation, userMessage)
	client := openai.NewClient(config.GetChatGptToken())
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: conversation,
		},
	)
	if err != nil {
		msg := fmt.Sprintf("ChatCompletion error: %s", err)
		slog.Error(msg)
		return msg
	}
	response := resp.Choices[0].Message.Content
	appendToConversationFile([]openai.ChatCompletionMessage{userMessage, {
		Role:    openai.ChatMessageRoleAssistant,
		Content: response,
	}})
	return response
}

func getConversation() []openai.ChatCompletionMessage {
	data, err := os.ReadFile(config.GetChatGptConversationPath())
	if err != nil {
		slog.Warn("No ChatGPT conversation file found.", "err", err)
		return nil
	}

	var conversation []openai.ChatCompletionMessage
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var message openai.ChatCompletionMessage
		err := json.Unmarshal([]byte(line), &message)
		if err != nil {
			continue
		}
		conversation = append(conversation, message)
	}

	return conversation
}

func appendToConversationFile(conversation []openai.ChatCompletionMessage) {
	f, err := os.OpenFile(config.GetChatGptConversationPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("Unable to open/create ChatGPT conversation file.", "err", err)
		return
	}
	defer f.Close()

	for _, message := range conversation {
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			slog.Error("Unable to marshal ChatGPT message.", "err", err)
			return
		}
		_, err = f.Write(append(jsonMessage, '\n'))
		if err != nil {
			slog.Error("Could not write ChatGPT conversation file.", "err", err)
			return
		}
	}
}
