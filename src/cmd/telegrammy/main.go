package main

import (
	"log/slog"
	"telegrammy/internal/config"
	"telegrammy/internal/handlers"
	"telegrammy/internal/utils"
)

func main() {
	utils.InitializeLogger()
	slog.Info("Starting telegrammy.", "numberOfResponseJobs", len(config.GetResponseJobs()), "numberOfCronJobs", len(config.GetCronJobs()))

	handlers.SetUpCronJobs()
	handlers.StartPollingLoop()
}
