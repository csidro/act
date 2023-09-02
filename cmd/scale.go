package cmd

import (
	"context"
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/csidro/act/prompts"
	"github.com/csidro/act/types"
	"github.com/sirupsen/logrus"
)

type ScaleCmd struct {
	Cluster string `help:"Cluster name or ARN"`
	Service string `help:"Service to scale"`
}

func (cmd *ScaleCmd) Run(ctx *types.ActContext) error {
	client := ecs.NewFromConfig(ctx.AwsConfig)

	clustersResponse, err := client.ListClusters(context.TODO(), &ecs.ListClustersInput{})
	if err != nil {
		logrus.Fatalf("failed to list clusters, %v", err)
	}

	sort.Strings(clustersResponse.ClusterArns)
	idx, _, err := prompts.SelectAwsResource("Cluster", clustersResponse.ClusterArns, cmd.Cluster)
	if err != nil {
		logrus.Fatalf("failed to select cluster, %v", err)
	}

	cluster := clustersResponse.ClusterArns[idx]
	servicesResponse, err := client.ListServices(context.TODO(), &ecs.ListServicesInput{Cluster: &cluster})
	if err != nil {
		logrus.Fatalf("failed to list services, %v", err)
	}

	sort.Strings(servicesResponse.ServiceArns)
	idx, _, err = prompts.SelectAwsResource("Service", servicesResponse.ServiceArns, cmd.Service)
	if err != nil {
		logrus.Fatalf("failed to select service, %v", err)
	}

	service := servicesResponse.ServiceArns[idx]
	fmt.Println(service)

	return nil
}
