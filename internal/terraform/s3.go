package terraform

type S3Config struct {
	Name       string       `yaml:"name"`
	Properties S3Properties `yaml:"properties"`
}

type S3Properties struct {
	BucketName   string            `yaml:"bucketName"`
	Region       string            `yaml:"region"`
	PublicAccess bool              `yaml:"publicAccess"`
	Tags         map[string]string `yaml:"tags"`
}
