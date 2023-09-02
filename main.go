package main

import (
	"github.com/csidro/act/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	cmd.Run()
}
