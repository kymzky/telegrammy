package domain

import (
	"fmt"
	"strings"
	"telegrammy/internal/utils"
)

type Job struct {
	Message          string   `yaml:"message,omitempty"`
	ParseMode        string   `yaml:"parseMode,omitempty"`
	Command          string   `yaml:"command"`
	EscapeCharacters []string `yaml:"escapeCharacters"`
}

func (job *Job) GetMessage(shellPath string) *string {
	var message string
	if job.Command == "" {
		message = job.Message
	} else {
		executor := utils.NewExecutor()
		output := executor.Execute(shellPath, job.Command)
		if len(job.EscapeCharacters) > 0 {
			output = job.escapeCharacters(output)
		}
		if strings.Contains(job.Message, "%s") {
			message = fmt.Sprintf(job.Message, output)
		} else {
			message = job.Message
		}
	}
	return &message
}

func (job *Job) escapeCharacters(text string) string {
	chars := job.EscapeCharacters
	for _, char := range chars {
		text = strings.ReplaceAll(text, char, "\\"+char)
	}
	return text
}
