package aws_handlers

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go/service/cloudtrail"
)

func GetCloudTrailEvents(cloudTrailClient *cloudtrail.CloudTrail) ([]map[string]interface{}, error) {
	input := &cloudtrail.LookupEventsInput{}

	events, err := cloudTrailClient.LookupEvents(input)

	if err != nil {
		return nil, errors.New("Error while fetching AWS Cloud Trail Events")
	}

	data := make([]map[string]interface{}, len(events.Events))

	for i, event := range events.Events {

		eventData := make(map[string]interface{})

		var cloudTrailEventData map[string]interface{}
		json.Unmarshal([]byte(*event.CloudTrailEvent), &cloudTrailEventData)

		eventData["EventId"] = *event.EventId
		eventData["EventName"] = *event.EventName
		eventData["EventTime"] = *event.EventTime
		eventData["EventSource"] = *event.EventSource
		eventData["EventType"] = cloudTrailEventData["eventType"]
		eventData["Region"] = cloudTrailEventData["awsRegion"]
		eventData["username"] = "-"

		if event.Username != nil {
			eventData["username"] = *event.Username
		}

		data[i] = eventData
	}

	return data, nil

}
