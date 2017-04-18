package awsresource

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/opsworks"
)

func (m *mockOpsWorksClient) DescribeApps(input *opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error) {
	resp := `
{
   "Apps": [
      {
         "AppId": "string",
         "AppSource": {
            "Password": "string",
            "Revision": "string",
            "SshKey": "string",
            "Type": "string",
            "Url": "string",
            "Username": "string"
         },
         "Attributes": {
            "string" : "string"
         },
         "CreatedAt": "string",
         "DataSources": [
            {
               "Arn": "string",
               "DatabaseName": "string",
               "Type": "string"
            }
         ],
         "Description": "string",
         "Domains": [ "string" ],
         "EnableSsl": "boolean",
         "Environment": [
            {
               "Key": "string",
               "Secure": "boolean",
               "Value": "string"
            }
         ],
         "Name": "string",
         "Shortname": "string",
         "SslConfiguration": {
            "Certificate": "string",
            "Chain": "string",
            "PrivateKey": "string"
         },
         "StackId": "string",
         "Type": "string"
      }
   ]
}`

	out := &opsworks.DescribeAppsOutput{}
	json.Unmarshal([]byte(resp), out)
	return out, nil
}

func TestDescribeOpsWorksApplication(t *testing.T) {
	mockSvc := &mockOpsWorksClient{}
	out, _ := DescribeOpsWorksApplication(mockSvc)
	got := strings.Join(out, "\n")
	want := `resource  "aws_opsworks_application" "string" {
  "id": "string",
  "name": "string",
  "short_name": "string",
  "stack_id": "string",
  "type": "string",
  "description": "string",
  "domains": [
    "string"
  ],
  "environment": [
    {
      "key": "string",
      "value": "string",
      "secure": false
    }
  ],
  "app_source": {
    "type": "string",
    "url": "string",
    "username": "string",
    "password": "string",
    "ssh_key": "string",
    "revision": "string"
  },
  "data_source": [
    {
      "data_source_arn": "string",
      "data_source_type": "string",
      "data_source_database_name": "string"
    }
  ],
  "enable_ssl": false,
  "ssl_configuration": {
    "private_key": "string",
    "certificate": "string",
    "chain": "string"
  },
  "document_root": null,
  "auto_bundle_on_deploy": null,
  "rails_env": null,
  "aws_flow_ruby_settings": null
}`

	if got != want {
		t.Errorf("DescribeOpsWorksApplication(svc)\ngot=%s,\nwant=%q", got, want)
	}
}
