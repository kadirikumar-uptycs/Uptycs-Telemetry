package gcp_handlers

import (
	"errors"
	"fmt"

	"github.com/1722101709/Telemetry/pkg/handlers"
	"google.golang.org/api/compute/v1"
)

func GetComputeInfo(service *compute.Service, projectId string) ([]map[string]interface{}, error) {

	// zones, err := service.Zones.List(projectId).Do()

	// if err != nil {
	// 	return nil, errors.New("Error while retrieving zones")
	// }

	zones := []string{"us-central1-a", "asia-south1-c"}

	data := make([]map[string]interface{}, 0)

	for _, zone := range zones {

		ZONE := handlers.GetLastString(zone, "/")

		response, err := service.Instances.List(projectId, ZONE).Do()

		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error while retrieving instances in %s zone", zone))
		}

		for _, instance := range response.Items {
			instanceData := make(map[string]interface{})

			instanceData["instanceName"] = instance.Name
			instanceData["instanceZone"] = ZONE
			instanceData["instanceType"] = handlers.GetLastString(instance.MachineType, "/")
			instanceData["instanceCPUPlatform"] = instance.CpuPlatform
			instanceData["instanceExternalIPAddress"] = instance.NetworkInterfaces[0].AccessConfigs[0].NatIP
			instanceData["instanceCreatedTime"] = instance.CreationTimestamp

			tags := make([]map[string]interface{}, 0)

			for key, value := range instance.Labels {
				tags = append(tags, map[string]interface{}{
					"key":   key,
					"value": value,
				})
			}
			instanceData["instanceTags"] = tags

			data = append(data, instanceData)
		}

	}
	return data, nil
}
