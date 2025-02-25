package domain

type ResponseJob struct {
	Trigger string `yaml:"trigger,omitempty"`
	Job     `yaml:",inline"`
}
