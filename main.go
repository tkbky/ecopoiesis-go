package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/opsworks"
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

				resp, err := svc.DescribeStacks(nil)

				if err != nil {
					fmt.Println("Fail to describe stacks", err)
					return err
				}

				var stacks []awsresource.OpsWorksStack
				for _, s := range resp.Stacks {
					stacks = append(stacks, awsresource.NewOpsWorksStack(s))
				}

				for _, s := range stacks {
					resource := Resource{Kind: "aws_opsworks_stack", Name: *s.Name, Obj: s}
					fmt.Printf("%s\n", resource.tf())
				}

				return nil
			},
		}, {
			Name:  "aocl",
			Usage: "OpsWorks Custom Layer",
			Action: func(c *cli.Context) error {
				svc := awsresource.NewOpsWorksService()

				resp, err := svc.DescribeStacks(nil)

				if err != nil {
					fmt.Println("Fail to describe stacks", err)
					return err
				}

				var layers []awsresource.OpsWorksCustomLayer
				for _, s := range resp.Stacks {
					resp, err := svc.DescribeLayers(&opsworks.DescribeLayersInput{StackId: aws.String(*s.StackId)})

					if err != nil {
						fmt.Printf("Fail to describe layer for stack `%s`: %s", *s.StackId, err)
						return err
					}

					for _, l := range resp.Layers {
						layers = append(layers, awsresource.NewOpsWorksCustomLayer(l))
					}
				}

				for _, l := range layers {
					resource := Resource{Kind: "aws_opsworks_custom_layer", Name: *l.Name, Obj: l}
					fmt.Printf("%s\n", resource.tf())
				}

				return nil
			},
		}, {
			Name:  "aoa",
			Usage: "OpsWorks Application",
			Action: func(c *cli.Context) error {
				svc := awsresource.NewOpsWorksService()

				resp, err := svc.DescribeStacks(nil)

				if err != nil {
					fmt.Println("Fail to describe application", err)
					return err
				}

				var apps []awsresource.OpsWorksApplication
				for _, s := range resp.Stacks {
					resp, err := svc.DescribeApps(&opsworks.DescribeAppsInput{StackId: aws.String(*s.StackId)})

					if err != nil {
						fmt.Printf("Fail to describe layer for stack `%s`: %s", *s.StackId, err)
						return err
					}

					for _, a := range resp.Apps {
						apps = append(apps, awsresource.NewOpsWorksApplication(a))
					}
				}

				for _, a := range apps {
					resource := Resource{Kind: "aws_opsworks_application", Name: *a.Name, Obj: a}
					fmt.Printf("%s\n", resource.tf())
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
