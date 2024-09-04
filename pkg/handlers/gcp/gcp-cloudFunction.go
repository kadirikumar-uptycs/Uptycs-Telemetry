package gcp_handlers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/1722101709/Telemetry/pkg/handlers"
	"google.golang.org/api/cloudfunctions/v1"
)

func GetCloudFunctionsInfo(service *cloudfunctions.Service, projectId string) ([]map[string]interface{}, error) {

	locations, err := service.Projects.Locations.List(fmt.Sprintf("projects/%s", projectId)).Do()

	if err != nil {
		return nil, errors.New("Error while fetching all Locations")
	}

	var data []map[string]interface{}

	for _, location := range locations.Locations {
		functions, err := service.Projects.Locations.Functions.List(location.Name).Do()

		if err != nil {
			return nil, errors.New("Error while retrieving functions")
		}
		for _, function := range functions.Functions {
			functionData := make(map[string]interface{})
			functionData["functionName"] = handlers.GetLastString(function.Name, "/")
			functionData["functionId"] = function.Name
			functionData["functionTrigger"] = function.HttpsTrigger.Url
			functionData["functionRegion"] = strings.Split(location.Name, "/")[3]
			functionData["functionRunTime"] = function.Runtime
			functionData["functionMemoryAllocated"] = strconv.FormatInt(function.AvailableMemoryMb, 10) + "MB"

			tags := make([]map[string]interface{}, 0)

			for key, value := range function.Labels {
				tags = append(tags, map[string]interface{}{
					"key":   key,
					"value": value,
				})
			}

			functionData["functionLabels"] = tags

			data = append(data, functionData)
		}
	}
	if len(data) == 0 {
		return make([]map[string]interface{}, 0), nil
	}
	return data, nil
}
