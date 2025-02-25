package config

import (
	"log/slog"
	"os"
	"sync"
	"telegrammy/internal/domain"

	"gopkg.in/yaml.v3"
)

var (
	config     *domain.Config
	configOnce sync.Once
)

func GetShellPath() string {
	config := getConfig()
	return config.ShellPath
}

func GetPollInterval() int {
	config := getConfig()
	return config.PollInterval
}

func GetResponseJobs() []domain.ResponseJob {
	config := getConfig()
	return config.ResponseJobs
}

func GetCronJobs() []domain.CronJob {
	config := getConfig()
	return config.CronJobs
}

func getConfig() *domain.Config {
	configOnce.Do(func() {
		configPath := GetTelegrammyConfigPath()
		data, err := os.ReadFile(configPath)
		if err != nil {
			slog.Error("Failed to read config file.", "err", err, "configPath", configPath)
			return
		}

		var cfg domain.Config
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			slog.Error("Error parsing config file", "err", err, "configPath", configPath)
			return
		}

		slog.Info("Configuration loaded successfully.")
		config = &cfg
	})
	return config
}
