package handlers

import (
	"log/slog"
	"telegrammy/internal/config"
	"telegrammy/internal/domain"
	"telegrammy/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)

func SetUpCronJobs() {
	cronJobs := config.GetCronJobs()
	for _, cronJob := range cronJobs {
		setUpCronJob(cronJob)
	}
}

func setUpCronJob(cronJob domain.CronJob) {
	cron := cron.New()
	cron.AddFunc(cronJob.Schedule, func() {
		message := tgbotapi.NewMessage(config.GetTelegramChatId(), *cronJob.GetMessage(config.GetShellPath()))
		message.ParseMode = cronJob.ParseMode
		slog.Info("Sending scheduled message.", "content", message.Text)
		telegram.SendMessage(&message)
	})
	cron.Start()
}
