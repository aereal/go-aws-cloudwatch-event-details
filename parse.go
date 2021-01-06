package cweventdetails

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

var (
	DetailTypeECSTaskStateChange              = "ECS Task State Change"
	DetailTypeECSContainerInstanceStateChange = "ECS Container Instance State Change"
	DetailTypeECSServiceDeploymentStateChange = "ECS Deployment State Change"
)

func ParseEventDetail(ev events.CloudWatchEvent) (interface{}, error) {
	var payload interface{}
	switch ev.DetailType {
	case DetailTypeECSTaskStateChange:
		payload = &ECSTaskStateChangeEvent{}
	case DetailTypeECSContainerInstanceStateChange:
		payload = &ECSContainerInstanceStateChangeEvent{}
	case DetailTypeECSServiceDeploymentStateChange:
		payload = &ECSDeploymentStateChangeEvent{}
	}
	if err := json.Unmarshal(ev.Detail, &payload); err != nil {
		return nil, err
	}
	return payload, nil
}
