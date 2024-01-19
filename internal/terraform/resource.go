package terraform

type ResourceConfig struct {
	EC2 *EC2Config `yaml:"ec2"`
	S3  *S3Config  `yaml:"s3"`
	SQS *SQSConfig `yaml:"sqs"`
}
