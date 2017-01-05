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

type resource interface {
	tf() string
}

// Resource describes a terraform resource
type Resource struct {
	Kind string
	Name string
	Obj  interface{}
}

func (r *Resource) tf() string {
	body, _ := json.MarshalIndent(r.Obj, "", "  ")
	return fmt.Sprintf("resource  \"%s\" \"%s\" %s", r.Kind, r.Name, body)
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

// OpsWorksStack describe an opsworks stack
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

// OpsWorksCustomLayer describes an opsworks custom layer
type OpsWorksCustomLayer struct {
	ID                       *string    `json:"id"`
	Name                     *string    `json:"name"`
	ShortName                *string    `json:"short_name"`
	StackID                  *string    `json:"stack_id"`
	AutoAssignElasticIPs     *bool      `json:"auto_assign_elastic_ips"`
	AutoAssignPublicIPs      *bool      `json:"auto_assign_public_ips"`
	CustomInstanceProfileArn *string    `json:"custom_instance_profile_arn"`
	CustomSecurityGroupIDs   []*string  `json:"custom_security_group_ids"`
	AutoHealing              *bool      `json:"auto_healing"`
	InstallUpdatesOnBoot     *bool      `json:"install_updates_on_boot"`
	ElasticLoadBalancer      *bool      `json:"elastic_load_balancer"`
	DrainELBOnShutdown       *bool      `json:"drain_elb_on_shutdown"`
	SystemPackages           []*string  `json:"system_packages"`
	UseEBSOptimizedInstances *bool      `json:"use_ebs_optimized_instances"`
	EBSVolume                *EBSVolume `json:"ebs_volume"`
	CustomJSON               *string    `json:"custom_json"`
	CustomConfigureRecipes   []*string  `json:"custom_configure_recipes"`
	CustomDeployRecipes      []*string  `json:"custom_deploy_recipes"`
	CustomSetupRecipes       []*string  `json:"custom_setup_recipes"`
	CustomShutdownRecipes    []*string  `json:"custom_shutdown_recipes"`
	CustomUndeployRecipes    []*string  `json:"custom_undeploy_recipes"`
}

// EBSVolume describes an ebs volume
type EBSVolume struct {
	MountPoint    *string `json:"mount_point"`
	Size          *string `json:"size"`
	NumberOfDisks *int    `json:"number_of_disks"`
	RAIDLevel     *int    `json:"raid_level"`
	VolumeType    *string `json:"type"`
	IOPS          *int    `json:"iops"`
}

func newOpsWorksService() *opsworks.OpsWorks {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("ap-southeast-1")})

	if err != nil {
		panic(err)
	}

	return opsworks.New(sess)
}

func newOpsWorksStack(s *opsworks.Stack) OpsWorksStack {
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

	return stack
}

func newOpsWorksCustomLayer(l *opsworks.Layer) OpsWorksCustomLayer {
	layer := OpsWorksCustomLayer{
		ID:                       l.LayerId,
		AutoAssignElasticIPs:     l.AutoAssignElasticIps,
		AutoAssignPublicIPs:      l.AutoAssignPublicIps,
		CustomInstanceProfileArn: l.CustomInstanceProfileArn,
		CustomJSON:               l.CustomJson,
		CustomSecurityGroupIDs:   l.CustomSecurityGroupIds,
		AutoHealing:              l.EnableAutoHealing,
		InstallUpdatesOnBoot:     l.InstallUpdatesOnBoot,
		Name:                     l.Name,
		SystemPackages:           l.Packages,
		ShortName:                l.Shortname,
		StackID:                  l.StackId,
		UseEBSOptimizedInstances: l.UseEbsOptimizedInstances,
		CustomConfigureRecipes:   l.CustomRecipes.Configure,
		CustomDeployRecipes:      l.CustomRecipes.Deploy,
		CustomSetupRecipes:       l.CustomRecipes.Setup,
		CustomShutdownRecipes:    l.CustomRecipes.Shutdown,
		CustomUndeployRecipes:    l.CustomRecipes.Undeploy,
	}

	return layer
}

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
				svc := newOpsWorksService()

				resp, err := svc.DescribeStacks(nil)

				if err != nil {
					fmt.Println("Fail to describe stacks", err)
					return err
				}

				var stacks []OpsWorksStack
				for _, s := range resp.Stacks {
					stacks = append(stacks, newOpsWorksStack(s))
				}

				for _, s := range stacks {
					resource := Resource{Kind: "aws_opsworks_stack", Name: *s.Name, Obj: s}
					fmt.Printf("%s\n", resource.tf())
				}

				return nil
			},
		}, {
			Name:  "aocl",
			Usage: "Opsworks Custom Layer",
			Action: func(c *cli.Context) error {
				svc := newOpsWorksService()

				resp, err := svc.DescribeStacks(nil)

				if err != nil {
					fmt.Println("Fail to describe stacks", err)
					return err
				}

				var layers []OpsWorksCustomLayer
				for _, s := range resp.Stacks {
					resp, err := svc.DescribeLayers(&opsworks.DescribeLayersInput{StackId: aws.String(*s.StackId)})

					if err != nil {
						fmt.Printf("Fail to describe layer for stack `%s`: %s", s.StackId, err)
						return err
					}

					for _, l := range resp.Layers {
						layers = append(layers, newOpsWorksCustomLayer(l))
					}
				}

				for _, l := range layers {
					resource := Resource{Kind: "aws_opsworks_custom_layer", Name: *l.Name, Obj: l}
					fmt.Printf("%s\n", resource.tf())
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
