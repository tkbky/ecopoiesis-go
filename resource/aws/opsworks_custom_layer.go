package awsresource

import "github.com/aws/aws-sdk-go/service/opsworks"

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

// NewOpsWorksCustomLayer returns an opsworks.Layer
func NewOpsWorksCustomLayer(l *opsworks.Layer) OpsWorksCustomLayer {
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
