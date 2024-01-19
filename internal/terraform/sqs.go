package terraform

type SQSConfig struct {
	Name       string        `yaml:"name"`
	Properties SQSProperties `yaml:"properties"`
}

type SQSProperties struct {
	VisibilityTimeout             int               `yaml:"visibilityTimeout"`
	MessageRetentionPeriod        int               `yaml:"messageRetentionPeriod"`
	DelaySeconds                  int               `yaml:"delaySeconds"`
	ReceiveMessageWaitTimeSeconds int               `yaml:"receiveMessageWaitTimeSeconds"`
	FifoQueue                     bool              `yaml:"fifoQueue"`
	ContentBasedDeduplication     bool              `yaml:"contentBasedDeduplication"`
	Tags                          map[string]string `yaml:"tags"`
}
