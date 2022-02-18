package airbyte

var DefaultAirbyteConfig = Config{
	AirbyteServerURL:      "http://localhost:8001",
	DestinationBucketName: "airbyte-data",
}

type Config struct {
	AWSAccessKeyID     string `mapstructure:"awsAccessKeyID" env:"AWS_ACCESS_KEY_ID"`
	AWSAccessKeySecret string `mapstructure:"awsAccessKeySecret" env:"AWS_ACCESS_KEY_SECRET"`
	AWSRegion          string `mapstructure:"awsRegion" env:"AWS_REGION"`
	AWSS3Endpoint      string `mapstructure:"awsS3Endpoint" env:"AWS_S3_ENDPOINT"`

	AirbyteServerURL      string `mapstructure:"airbyteServerURL" env:"AIRBYTE_SERVER_URL"`
	DestinationBucketName string `mapstructure:"airbyteDestinationBucket" env:"AIRBYTE_DESTINATION_BUCKET_NAME"`
}
