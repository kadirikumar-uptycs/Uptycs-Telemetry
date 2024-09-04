package aws_handlers

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetVPCs(client *ec2.EC2) ([]map[string]interface{}, error) {

	vpcResponse, err := client.DescribeVpcs(&ec2.DescribeVpcsInput{})

	if err != nil {
		return nil, errors.New("Error While Fetching the VPCs")
	}

	data := make([]map[string]interface{}, len(vpcResponse.Vpcs))

	for i, vpc := range vpcResponse.Vpcs {
		vpcData := make(map[string]interface{})

		vpcData["VPCId"] = *vpc.VpcId
		vpcData["IsDefaultVPC"] = *vpc.IsDefault
		vpcData["VPCOwnerId"] = *vpc.OwnerId
		vpcData["VPCState"] = *vpc.State
		vpcData["VPCCidrBlock"] = *vpc.CidrBlock

		data[i] = vpcData
	}
	return data, nil
}
