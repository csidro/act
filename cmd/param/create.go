package param

import (
	"context"
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/csidro/act/prompts"
	"github.com/csidro/act/types"
	"github.com/sirupsen/logrus"
)

type CreateCmd struct {
	Path  string `arg:"path"`
	Value string `arg:"value"`

	Description string                 `arg:"" optional:""`
	Type        ssmTypes.ParameterType `enum:"String,StringList,SecureString" default:"String"`
	KeyId       string                 ``
	Overwrite   bool                   `default:"false"`

	DataType string                 `enum:"text" default:"text"`
	Tier     ssmTypes.ParameterTier `enum:"Standard" default:"Standard"`
}

func (cmd *CreateCmd) Run(ctx *types.ActContext) error {
	// Force supported defaults
	cmd.DataType = "text"
	cmd.Tier = ssmTypes.ParameterTierStandard

	client := ssm.NewFromConfig(ctx.AwsConfig)
	kmsClient := kms.NewFromConfig(ctx.AwsConfig)

	putParameterInput := &ssm.PutParameterInput{
		Name:        &cmd.Path,
		Value:       &cmd.Value,
		DataType:    &cmd.DataType,
		Description: &cmd.Description,
		Overwrite:   &cmd.Overwrite,
		Tier:        cmd.Tier,
		Type:        cmd.Type,
	}

	if cmd.Type == ssmTypes.ParameterTypeSecureString && cmd.KeyId == "" {
		aliasName, err := selectKmsKey(cmd, ctx, *kmsClient)
		if err != nil {
			logrus.Fatal(err)
		}

		putParameterInput.KeyId = aliasName
	}

	_, err := client.PutParameter(
		context.TODO(),
		putParameterInput,
	)

	if err != nil {
		logrus.Fatalf("failed to create parameter, %v", err)
	}

	return nil
}

func selectKmsKey(cmd *CreateCmd, ctx *types.ActContext, client kms.Client) (*string, error) {
	aliasesResponse, err := client.ListAliases(context.TODO(), &kms.ListAliasesInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to list kms key aliases, %v", err)
	}

	sort.Slice(aliasesResponse.Aliases, func(a, b int) bool {
		return *aliasesResponse.Aliases[a].AliasName < *aliasesResponse.Aliases[b].AliasName
	})

	idx, _, err := prompts.SelectKmsAlias(aliasesResponse.Aliases, "")
	if err != nil {
		return nil, fmt.Errorf("failed to select a key, %v", err)
	}

	return aliasesResponse.Aliases[idx].AliasName, nil
}
