package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/urfave/cli"
)

type awsresource interface {
	tf() string
}

// Resource describes a terraform resource
type Resource struct {
	Kind string
	Name string
	Obj  interface{}
}

// CustomCookbooksSource describes a custom cookbooks source
type CustomCookbooksSource struct {
	SourceType *string `json:"type"`
	URL        *string `json:"url"`
	Username   *string `json:"username"`
	Password   *string `json:"password"`
	SSHKey     *string `json:"ssh_key"`
	Revision   *string `json:"revision"`
}

func (r *Resource) tf() string {
	body, _ := json.MarshalIndent(r.Obj, "", "  ")
	return fmt.Sprintf("resource  \"%s\" \"%s\" %s", r.Kind, r.Name, body)
}

// OpsWorksStack describe  an opsworks stack
type OpsWorksStack struct {
	ID                          *string                `json:"id"`
	Name                        *string                `json:"name"`
	Region                      *string                `json:"region"`
	ServiceRoleArn              *string                `json:"service_role_arn"`
	DefaultInstanceProfileArn   *string                `json:"default_instance_profile_arn"`
	AgentVersion                *string                `json:"agent_version"`
	BerkshelfVersion            *string                `json:"berkshelf_version"`
	Color                       *string                `json:"color"`
	DefaultAvailabilityZone     *string                `json:"default_availability_zone"`
	ConfigurationManagerName    *string                `json:"configuration_manager_name"`
	ConfigurationManagerVersion *string                `json:"configuration_manager_version"`
	CustomCookbooksSource       *CustomCookbooksSource `json:"custom_cookbooks_source"`
	CustomJSON                  *string                `json:"custom_json"`
	DefaultOS                   *string                `json:"default_os"`
	DefaultRootDeviceType       *string                `json:"default_root_device_type"`
	DefaultSSHKeyName           *string                `json:"default_ssh_key_name"`
	DefaultSubnetID             *string                `json:"default_subnet_id"`
	HostnameTheme               *string                `json:"hostname_theme"`
	ManageBerkshelf             *bool                  `json:"manage_berkshelf"`
	UseCustomCookbooks          *bool                  `json:"use_custom_cookbooks"`
	UseOpsworksSecurityGroups   *bool                  `json:"use_opsworks_security_groups"`
	VPCID                       *string                `json:"vpc_id"`
}

func main() {
	app := cli.NewApp()
	app.Name = "trf"
	app.Usage = "Export AWS resources to Terraform resources"

	app.Commands = []cli.Command{
		{
			Name:  "aocl",
			Usage: "OpsWorks Custom Layer Resource",
			Action: func(c *cli.Context) error {
				sess, err := session.NewSession(&aws.Config{Region: aws.String("ap-southeast-1")})

				if err != nil {
					fmt.Println("Fail to create session", err)
					return err
				}

				svc := opsworks.New(sess)

				resp, err := svc.DescribeStacks(nil)

				if err != nil {
					fmt.Println("Fail to describe stacks", err)
					return err
				}

				var stacks []OpsWorksStack
				for _, s := range resp.Stacks {
					stack := OpsWorksStack{
						ID:                        s.StackId,
						Name:                      s.Name,
						Region:                    s.Region,
						ServiceRoleArn:            s.ServiceRoleArn,
						DefaultInstanceProfileArn: s.DefaultInstanceProfileArn,
						AgentVersion:              s.AgentVersion,
						BerkshelfVersion:          s.ChefConfiguration.BerkshelfVersion,
						Color:                     s.Attributes["Color"],
						DefaultAvailabilityZone:     s.DefaultAvailabilityZone,
						ConfigurationManagerName:    s.ConfigurationManager.Name,
						ConfigurationManagerVersion: s.ConfigurationManager.Version,
						CustomCookbooksSource: &CustomCookbooksSource{
							SourceType: s.CustomCookbooksSource.Type,
							URL:        s.CustomCookbooksSource.Url,
							Username:   s.CustomCookbooksSource.Username,
							Password:   s.CustomCookbooksSource.Password,
							SSHKey:     s.CustomCookbooksSource.SshKey,
							Revision:   s.CustomCookbooksSource.Revision,
						},
						CustomJSON:                s.CustomJson,
						DefaultOS:                 s.DefaultOs,
						DefaultRootDeviceType:     s.DefaultRootDeviceType,
						DefaultSSHKeyName:         s.DefaultSshKeyName,
						DefaultSubnetID:           s.DefaultSubnetId,
						HostnameTheme:             s.HostnameTheme,
						ManageBerkshelf:           s.ChefConfiguration.ManageBerkshelf,
						UseCustomCookbooks:        s.UseCustomCookbooks,
						UseOpsworksSecurityGroups: s.UseOpsworksSecurityGroups,
						VPCID: s.VpcId,
					}

					stacks = append(stacks, stack)
				}

				for _, s := range stacks {
					resource := Resource{Kind: "aws_opsworks_stack", Name: *s.Name, Obj: s}
					fmt.Printf("%s\n", resource.tf())
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
