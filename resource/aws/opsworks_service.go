package awsresource

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworks"
)

// NewOpsWorksService returns an opsworks.OpsWorks
func NewOpsWorksService() *opsworks.OpsWorks {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("ap-southeast-1")})

	if err != nil {
		panic(err)
	}

	return opsworks.New(sess)
}
