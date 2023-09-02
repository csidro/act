package templates

import (
	"strings"

	kmsTypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
	ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/samber/lo"
)

type AwsResource struct {
	Name string
	Arn  string
}

func AwsResources(list []string) []AwsResource {
	return lo.Map(list, func(item string, idx int) AwsResource {
		parts := strings.Split(item, "/")
		return AwsResource{Name: parts[len(parts)-1], Arn: item}
	})
}

func Parameters(list []ssmTypes.Parameter) []AwsResource {
	return lo.Map(list, func(item ssmTypes.Parameter, idx int) AwsResource {
		return AwsResource{Name: *item.Name, Arn: *item.ARN}
	})
}

func KmsAliases(list []kmsTypes.AliasListEntry) []AwsResource {
	return lo.Map(list, func(item kmsTypes.AliasListEntry, idx int) AwsResource {
		return AwsResource{Name: *item.AliasName, Arn: *item.AliasArn}
	})
}
