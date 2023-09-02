package param

import (
	"context"
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/csidro/act/prompts"
	"github.com/csidro/act/types"
	"github.com/sirupsen/logrus"
)

type ReadCmd struct {
	Path string `arg:"path" optional:""`

	RootPath       string `default:"/"`
	Recursive      bool   `default:"true"`
	WithDecryption bool   `default:"true"`
}

func (cmd *ReadCmd) Run(ctx *types.ActContext) error {
	// Force decryption
	cmd.WithDecryption = true

	client := ssm.NewFromConfig(ctx.AwsConfig)

	if cmd.Path == "" {
		paramsResponse, err := client.GetParametersByPath(
			context.TODO(),
			&ssm.GetParametersByPathInput{
				Path:           &cmd.RootPath,
				Recursive:      &cmd.Recursive,
				WithDecryption: &cmd.WithDecryption,
			},
		)

		if err != nil {
			logrus.Fatalf("failed to get parameters by path, %v", err)
		}

		sort.Slice(paramsResponse.Parameters, func(a, b int) bool {
			return *paramsResponse.Parameters[a].Name < *paramsResponse.Parameters[b].Name
		})

		idx, _, err := prompts.SelectParameter(paramsResponse.Parameters, "")
		if err != nil {
			logrus.Fatalf("failed to select a parameter, %v", err)
		}

		fmt.Println(*paramsResponse.Parameters[idx].Value)
		return nil
	}

	paramResponse, err := client.GetParameter(
		context.TODO(),
		&ssm.GetParameterInput{Name: &cmd.Path, WithDecryption: &cmd.WithDecryption},
	)
	if err != nil {
		logrus.Fatalf("failed to get parameter from store, %v", err)
	}

	fmt.Println(*paramResponse.Parameter.Value)
	return nil
}
