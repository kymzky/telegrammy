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

	executor := utils.NewExecutor()
	if err := executeInitCommand(executor); err != nil {
		return err
	}

	return startMainApplication()
}

func executeInitCommand(executor *utils.Executor) error {
	initCmd := config.GetInitCommand()
	if initCmd == "" || len(initCmd) == 0 {
		return nil
	}

	slog.Info("Executing init command", "command", initCmd)

	result, err := executor.Execute(initCmd)
	if err != nil {
		slog.Error("Failed to execute command", "error", err)
		return err
	}

	if result.Stdout != "" {
		slog.Info("Command output:", "stdout", result.Stdout)
	}
	if result.Stderr != "" {
		slog.Warn("Command stderr:", "stderr", result.Stderr)
	}

	slog.Info("Command completed successfully", "cmd", initCmd)
	return nil
}

func startMainApplication() error {
	slog.Info("Starting telegrammy",
		"numberOfResponseJobs", len(config.GetResponseJobs()),
		"numberOfCronJobs", len(config.GetCronJobs()))

	handlers.SetUpCronJobs()
	handlers.StartPollingLoop()
	return nil
}
