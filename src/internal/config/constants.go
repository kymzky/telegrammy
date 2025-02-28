package config

import "path/filepath"

const (
	telegrammyDir           = "/etc/telegrammy"
	configFile              = "config.yaml"
	chatGptConversationFile = "conversation"
)

func GetConfigPath() string {
	return filepath.Join(telegrammyDir, configFile)
}

func GetChatGptConversationPath() string {
	return filepath.Join(telegrammyDir, chatGptConversationFile)
}
