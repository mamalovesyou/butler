package destinations

const S3_DESTINATION_CONFIG = `{
  "destinationDefinitionId":"4816b78f-1489-44c1-9060-4b19d5fa9362",
  "connectionConfiguration":{
    "secret_access_key":"AbCdEfGhEXAMPLEKEY",
    "s3_bucket_region":"us-east-2",
    "s3_bucket_path":"workspaces/<workspace_id>",
    "s3_bucket_name":"airbyte-data",
    "access_key_id":"A012345678910EXAMPLE",
    "s3_endpoint":"http://localhost:9000",
    "format":{
      "part_size_mb":5,
      "format_type":"JSONL"
    }
  }
}`

type S3DestinationFormatConfig struct {
	PartSizeMb int8   `json:"part_size_mb"`
	FormatType string `json:"format_type"`
}

type S3DestinationConfig struct {
	AccessKeyID     string                    `json:"access_key_id"`
	AccessKeySecret string                    `json:"secret_access_key"`
	S3BucketRegion  string                    `json:"s3_bucket_region"`
	S3Endpoint      string                    `json:"s3_endpoint"`
	S3BucketPath    string                    `json:"s3_bucket_path"`
	S3BucketName    string                    `json:"s3_bucket_name"`
	Format          S3DestinationFormatConfig `json:"format"`
}

func NewS3DestinationConfig(name, region, endpoint, accessKey, secretKey string) S3DestinationConfig {
	return S3DestinationConfig{
		Format: S3DestinationFormatConfig{
			PartSizeMb: 10,
			FormatType: "JSONL",
		},
		S3BucketName:    name,
		S3Endpoint:      endpoint,
		S3BucketRegion:  region,
		AccessKeyID:     accessKey,
		AccessKeySecret: secretKey,
	}
}

type S3Destination struct {
	DestinationDefinitionID string
	BaseConfig              S3DestinationConfig
}

func (dest *S3Destination) Name() string {
	return "S3"
}

func (dest *S3Destination) AirbyteDefinitionID() string {
	return dest.DestinationDefinitionID
}

func (dest *S3Destination) BindAirbyteDefinition(destinationDefinitionID string) {
	dest.DestinationDefinitionID = destinationDefinitionID
}

func (dest *S3Destination) BuildConfig(workspaceID string) interface{} {
	cfg := dest.BaseConfig
	cfg.S3BucketPath = workspaceID
	return cfg
}
