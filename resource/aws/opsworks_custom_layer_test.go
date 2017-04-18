package awsresource

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/opsworks"
)

func (m *mockOpsWorksClient) DescribeLayers(input *opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error) {
	resp := `
{
   "Layers": [
      {
         "Attributes": {
            "string" : "string"
         },
         "AutoAssignElasticIps": "boolean",
         "AutoAssignPublicIps": "boolean",
         "CloudWatchLogsConfiguration": {
            "Enabled": "boolean",
            "LogStreams": [
               {
                  "BatchCount": "number",
                  "BatchSize": "number",
                  "BufferDuration": "number",
                  "DatetimeFormat": "string",
                  "Encoding": "string",
                  "File": "string",
                  "FileFingerprintLines": "string",
                  "InitialPosition": "string",
                  "LogGroupName": "string",
                  "MultiLineStartPattern": "string",
                  "TimeZone": "string"
               }
            ]
         },
         "CreatedAt": "string",
         "CustomInstanceProfileArn": "string",
         "CustomJson": "string",
         "CustomRecipes": {
            "Configure": [ "string" ],
            "Deploy": [ "string" ],
            "Setup": [ "string" ],
            "Shutdown": [ "string" ],
            "Undeploy": [ "string" ]
         },
         "CustomSecurityGroupIds": [ "string" ],
         "DefaultRecipes": {
            "Configure": [ "string" ],
            "Deploy": [ "string" ],
            "Setup": [ "string" ],
            "Shutdown": [ "string" ],
            "Undeploy": [ "string" ]
         },
         "DefaultSecurityGroupNames": [ "string" ],
         "EnableAutoHealing": "boolean",
         "InstallUpdatesOnBoot": "boolean",
         "LayerId": "string",
         "LifecycleEventConfiguration": {
            "Shutdown": {
               "DelayUntilElbConnectionsDrained": "boolean",
               "ExecutionTimeout": "number"
            }
         },
         "Name": "string",
         "Packages": [ "string" ],
         "Shortname": "string",
         "StackId": "string",
         "Type": "string",
         "UseEbsOptimizedInstances": "boolean",
         "VolumeConfigurations": [
            {
               "Iops": "number",
               "MountPoint": "string",
               "NumberOfDisks": "number",
               "RaidLevel": "number",
               "Size": "number",
               "VolumeType": "string"
            }
         ]
      }
   ]
}`

	out := &opsworks.DescribeLayersOutput{}
	json.Unmarshal([]byte(resp), out)
	return out, nil
}

func TestDescribeOpsWorksCustomLayer(t *testing.T) {
	mockSvc := &mockOpsWorksClient{}
	out, _ := DescribeOpsWorksCustomLayer(mockSvc)
	got := strings.Join(out, "\n")
	want := `resource  "aws_opsworks_custom_layer" "string" {
  "id": "string",
  "name": "string",
  "short_name": "string",
  "stack_id": "string",
  "auto_assign_elastic_ips": false,
  "auto_assign_public_ips": false,
  "custom_instance_profile_arn": "string",
  "custom_security_group_ids": [
    "string"
  ],
  "auto_healing": false,
  "install_updates_on_boot": false,
  "elastic_load_balancer": null,
  "drain_elb_on_shutdown": null,
  "system_packages": [
    "string"
  ],
  "use_ebs_optimized_instances": false,
  "ebs_volume": null,
  "custom_json": "string",
  "custom_configure_recipes": [
    "string"
  ],
  "custom_deploy_recipes": [
    "string"
  ],
  "custom_setup_recipes": [
    "string"
  ],
  "custom_shutdown_recipes": [
    "string"
  ],
  "custom_undeploy_recipes": [
    "string"
  ]
}`

	if got != want {
		t.Errorf("DescribeOpsWorksCustomLayer(svc)\ngot=%s,\nwant=%q", got, want)
	}
}
