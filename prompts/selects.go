package prompts

import (
	"fmt"
	"strings"

	kmsTypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
	ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/csidro/act/prompts/templates"
	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
)

func SelectAwsResource(resourceType string, resources []string, name string) (int, string, error) {
	if len(resources) == 0 {
		msg := templates.PromptFailure(fmt.Sprintf("%s not found", resourceType), fmt.Sprintf("No available %s", resourceType))
		return -1, "", fmt.Errorf(msg)
	}

	resource, idx, found := lo.FindIndexOf(resources, func(item string) bool {
		parts := strings.Split(item, "/")
		return parts[len(parts)-1] == name
	})

	if name != "" && !found {
		msg := templates.PromptFailure(resourceType, name)
		fmt.Println(msg)
	}

	if found {
		msg := templates.PromptSuccess(resourceType, resource)
		fmt.Println(msg)
		return idx, resource, nil
	}

	prompt := promptui.Select{
		Label:     fmt.Sprintf("Select %s", resourceType),
		Items:     templates.AwsResources(resources),
		Templates: templates.TplAwsResource(resourceType),
	}

	return prompt.Run()
}

func SelectParameter(resources []ssmTypes.Parameter, name string) (int, string, error) {
	if len(resources) == 0 {
		msg := templates.PromptFailure("Parameter not found", "No available Parameter")
		return -1, "", fmt.Errorf(msg)
	}

	parameter, idx, found := lo.FindIndexOf(resources, func(param ssmTypes.Parameter) bool {
		return *param.Name == name
	})

	if name != "" && !found {
		msg := templates.PromptFailure("Parameter", name)
		fmt.Println(msg)
	}

	if found {
		msg := templates.PromptSuccess("Parameter", *parameter.Name)
		fmt.Println(msg)
		return idx, *parameter.Name, nil
	}

	prompt := promptui.Select{
		Label:     "Select Parameter",
		Items:     templates.Parameters(resources),
		Templates: templates.TplAwsResource("Parameter"),
	}

	return prompt.Run()
}

func SelectKmsAlias(resources []kmsTypes.AliasListEntry, name string) (int, string, error) {
	if len(resources) == 0 {
		msg := templates.PromptFailure("Key not found", "No available Key")
		return -1, "", fmt.Errorf(msg)
	}

	alias, idx, found := lo.FindIndexOf(resources, func(param kmsTypes.AliasListEntry) bool {
		return *param.AliasName == name
	})

	if name != "" && !found {
		msg := templates.PromptFailure("Key", name)
		fmt.Println(msg)
	}

	if found {
		msg := templates.PromptSuccess("Key", *alias.AliasName)
		fmt.Println(msg)
		return idx, *alias.AliasName, nil
	}

	prompt := promptui.Select{
		Label:     "Select Parameter",
		Items:     templates.KmsAliases(resources),
		Templates: templates.TplAwsResource("Kms key"),
	}

	return prompt.Run()
}
