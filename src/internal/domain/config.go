package domain

type Config struct {
	ShellPath    string        `yaml:"shellPath,omitempty"`
	ResponseJobs []ResponseJob `yaml:"responseJobs"`
	CronJobs []CronJob `yaml:"cronJobs"`
	PollInterval int           `yaml:"pollInterval,omitempty"`
}
