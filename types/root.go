package types

import "github.com/aws/aws-sdk-go-v2/aws"

type ActContext struct {
	Debug     bool
	AwsConfig aws.Config
}
