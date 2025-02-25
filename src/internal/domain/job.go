package domain

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
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
		output := executeCommand(shellPath, job.Command)
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

func executeCommand(shellPath string, command string) string {
	output, err := exec.Command(shellPath, "-c", command).CombinedOutput()
	if err != nil {
		msg := fmt.Sprintf("Error executing command '%s': %s\n\n%s", command, err, output)
		slog.Error(msg)
		return msg
	}
	return string(output)
}

func (job *Job) escapeCharacters(text string) string {
	chars := job.EscapeCharacters
	for _, char := range chars {
		text = strings.ReplaceAll(text, char, "\\"+char)
	}
	return text
}
