package cweventdetails

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

var (
	DetailTypeECSTaskStateChange = "ECS Task State Change"
)

func ParseEventDetail(ev events.CloudWatchEvent) (interface{}, error) {
	var payload interface{}
	switch ev.DetailType {
	case DetailTypeECSTaskStateChange:
		payload = &ECSTaskStateChangeEvent{}
	}
	if err := json.Unmarshal(ev.Detail, &payload); err != nil {
		return nil, err
	}
	return payload, nil
}
