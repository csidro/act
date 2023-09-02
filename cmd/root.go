package cmd

import (
	"context"

	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/csidro/act/cmd/param"
	"github.com/csidro/act/types"
	"github.com/sirupsen/logrus"
)

var Cli struct {
	Debug bool `help:"Enable debug mode."`
	Yolo  bool `help:"You only live once."`

	Scale ScaleCmd       `cmd:"scale" help:"Scale service."`
	Param param.ParamCmd `cmd:"param" help:"Paramstore stuff."`
}

func GetAwsConfig(ctx *types.ActContext) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logrus.Fatalf("failed to load configuration, %v", err)
	}

	ctx.AwsConfig = cfg
}

func Run() {
	ctx := kong.Parse(
		&Cli,
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}))

	cmdContext := &types.ActContext{Debug: Cli.Debug}
	GetAwsConfig(cmdContext)

	ctx.Run(cmdContext)
}
