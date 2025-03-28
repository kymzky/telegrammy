package main

import (
	"log/slog"
	"os"
	"telegrammy/internal/config"
	"telegrammy/internal/handlers"
	"telegrammy/internal/utils"
)

func main() {
	if err := runApplication(); err != nil {
		slog.Error("Application failed", "error", err)
		os.Exit(1)
	}
}

func runApplication() error {
	utils.InitializeLogger()
	executeInitCommand()
	return startMainApplication()
}

func executeInitCommand() {
	initCmd := config.GetInitCommand()
	if initCmd == "" || len(initCmd) == 0 {
		return
	}
	executor := utils.NewExecutor()
	executor.Execute(config.GetShellPath(), initCmd)
}

func startMainApplication() error {
	slog.Info("Starting telegrammy",
		"numberOfResponseJobs", len(config.GetResponseJobs()),
		"numberOfCronJobs", len(config.GetCronJobs()))

	handlers.SetUpCronJobs()
	handlers.StartPollingLoop()
	return nil
}
