package terraform

type EC2Config struct {
	Name       string        `yaml:"name"`
	Properties EC2Properties `yaml:"properties"`
}

type EC2Properties struct {
	InstanceType   string            `yaml:"instanceType"`
	AMI            string            `yaml:"ami"`
	KeyName        string            `yaml:"keyName"`
	SecurityGroups []string          `yaml:"securityGroups"`
	Tags           map[string]string `yaml:"tags"`
}
