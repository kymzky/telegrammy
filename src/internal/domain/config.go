package domain

type Config struct {
	ShellPath    string        `yaml:"shellPath,omitempty"`
	InitCommand  string        `yaml:"initCommand"`
	ResponseJobs []ResponseJob `yaml:"responseJobs"`
	CronJobs     []CronJob     `yaml:"cronJobs"`
	PollInterval int           `yaml:"pollInterval,omitempty"`
}
