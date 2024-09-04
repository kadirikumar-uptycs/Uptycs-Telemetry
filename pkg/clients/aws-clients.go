package clients

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	CloudTrailClient *cloudtrail.CloudTrail
	S3Client         *s3.S3
	EC2Client        *ec2.EC2
	IAMClient        *iam.IAM
)

func StartClients() {
	region := "us-east-1"
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		fmt.Printf("Error Creating New AWS Session")
		return
	}

	CloudTrailClient = cloudtrail.New(awsSession)
	S3Client = s3.New(awsSession)
	EC2Client = ec2.New(awsSession)
	IAMClient = iam.New(awsSession)
}
