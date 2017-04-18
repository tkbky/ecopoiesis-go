package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	awsresource "github.com/tkbky/trf/resource/aws"
)

func main() {
	app := cli.NewApp()
	app.Name = "trf"
	app.Usage = "Export AWS resources to Terraform resources"
	app.Version = Version

	app.Commands = []cli.Command{
		{
			Name:  "aos",
			Usage: "OpsWorks Stack",
			Action: func(c *cli.Context) error {
				svc := awsresource.NewOpsWorksService()
				tfs, err := awsresource.DescribeOpsWorksStack(svc)

				if err != nil {
					panic(err)
				}

				for _, s := range tfs {
					fmt.Printf("%s\n", s)
				}

				return nil
			},
		}, {
			Name:  "aocl",
			Usage: "OpsWorks Custom Layer",
			Action: func(c *cli.Context) error {
				svc := awsresource.NewOpsWorksService()
				tfs, err := awsresource.DescribeOpsWorksCustomLayer(svc)

				if err != nil {
					panic(err)
				}

				for _, s := range tfs {
					fmt.Printf("%s\n", s)
				}

				return nil
			},
		}, {
			Name:  "aoa",
			Usage: "OpsWorks Application",
			Action: func(c *cli.Context) error {
				svc := awsresource.NewOpsWorksService()
				tfs, err := awsresource.DescribeOpsWorksApplication(svc)

				if err != nil {
					panic(err)
				}

				for _, s := range tfs {
					fmt.Printf("%s\n", s)
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
