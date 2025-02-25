package domain

type CronJob struct {
	Schedule string `yaml:"schedule"`
	Job      `yaml:",inline"`
}
