package aws_handlers

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetEC2Instances(ec2Client *ec2.EC2) ([]map[string]string, error) {

	response, err := ec2Client.DescribeInstances(nil)

	if err != nil {
		return nil, errors.New("Error while fetching AWS EC2 Instances")
	}

	var data []map[string]string

	for _, reservation := range response.Reservations {
		for _, instance := range reservation.Instances {

			instanceName := ""
			var instanceId, publicIP, instanceType, keyName, platform, vpcId string

			if instance.InstanceId != nil {
				instanceId = *instance.InstanceId
			}

			if instance.PublicIpAddress != nil {
				publicIP = *instance.PublicIpAddress
			}

			if instance.InstanceType != nil {
				instanceType = *instance.InstanceType
			}

			if instance.KeyName != nil {
				keyName = *instance.KeyName
			}

			if instance.PlatformDetails != nil {
				platform = *instance.PlatformDetails
			}

			if instance.VpcId != nil {
				vpcId = *instance.VpcId
			}

			for _, tag := range instance.Tags {
				if tag.Key != nil && tag.Value != nil && *tag.Key == "Name" {
					instanceName = *tag.Value
					break
				}
			}

			instanceDetails := map[string]string{
				"InstanceName": instanceName,
				"PublicIp":     publicIP,
				"InstanceId":   instanceId,
				"InstanceType": instanceType,
				"keyName":      keyName,
				"platform":     platform,
				"vpcId":        vpcId,
			}

			data = append(data, instanceDetails)
		}
	}
	if len(data) == 0 {
		return make([]map[string]string, 0), nil
	}
	return data, nil
}
