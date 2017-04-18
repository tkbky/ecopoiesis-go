package awsresource

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/opsworks/opsworksiface"
)

type mockOpsWorksClient struct {
	opsworksiface.OpsWorksAPI
}

func (m *mockOpsWorksClient) DescribeStacks(*opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error) {
	resp := `
{
   "Stacks": [
      {
         "AgentVersion": "string",
         "Arn": "string",
         "Attributes": {
            "string" : "string"
         },
         "ChefConfiguration": {
            "BerkshelfVersion": "string",
            "ManageBerkshelf": "boolean"
         },
         "ConfigurationManager": {
            "Name": "string",
            "Version": "string"
         },
         "CreatedAt": "string",
         "CustomCookbooksSource": {
            "Password": "string",
            "Revision": "string",
            "SshKey": "string",
            "Type": "string",
            "Url": "string",
            "Username": "string"
         },
         "CustomJson": "string",
         "DefaultAvailabilityZone": "string",
         "DefaultInstanceProfileArn": "string",
         "DefaultOs": "string",
         "DefaultRootDeviceType": "string",
         "DefaultSshKeyName": "string",
         "DefaultSubnetId": "string",
         "HostnameTheme": "string",
         "Name": "string",
         "Region": "string",
         "ServiceRoleArn": "string",
         "StackId": "string",
         "UseCustomCookbooks": "boolean",
         "UseOpsworksSecurityGroups": "boolean",
         "VpcId": "string"
      }
   ]
}`

	out := &opsworks.DescribeStacksOutput{}
	json.Unmarshal([]byte(resp), out)
	return out, nil
}

func TestDescribeOpsWorksStack(t *testing.T) {
	mockSvc := &mockOpsWorksClient{}
	out, _ := DescribeOpsWorksStack(mockSvc)
	got := strings.Join(out, "\n")
	want := `resource  "aws_opsworks_stack" "string" {
  "id": "string",
  "name": "string",
  "region": "string",
  "service_role_arn": "string",
  "default_instance_profile_arn": "string",
  "agent_version": "string",
  "berkshelf_version": "string",
  "color": null,
  "default_availability_zone": "string",
  "configuration_manager_name": "string",
  "configuration_manager_version": "string",
  "custom_cookbooks_source": {
    "type": "string",
    "url": "string",
    "username": "string",
    "password": "string",
    "ssh_key": "string",
    "revision": "string"
  },
  "custom_json": "string",
  "default_os": "string",
  "default_root_device_type": "string",
  "default_ssh_key_name": "string",
  "default_subnet_id": "string",
  "hostname_theme": "string",
  "manage_berkshelf": false,
  "use_custom_cookbooks": false,
  "use_opsworks_security_groups": false,
  "vpc_id": "string"
}`

	if got != want {
		t.Errorf("DescribeOpsWorksStack(svc)\ngot=%q,\nwant=%q", got, want)
	}
}
