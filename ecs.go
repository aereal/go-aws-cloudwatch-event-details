package cweventdetails

import "time"

type TaskAttachmentDetail struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TaskAttachment struct {
	ID      string                 `json:"id"`
	Type    string                 `json:"type"`
	Status  string                 `json:"status"`
	Details []TaskAttachmentDetail `json:"details"`
}

type NetworkInterface struct {
	AttachmentID       string `json:"attachmentId"`
	PrivateIpv4Address string `json:"privateIpv4Address"`
}

// Container is a container in the task.
type Container struct {
	ContainerArn      string             `json:"containerArn"`
	LastStatus        string             `json:"lastStatus"`
	Name              string             `json:"name"`
	Image             string             `json:"image"`
	ImageDigest       string             `json:"imageDigest"`
	RuntimeID         string             `json:"runtimeId"`
	TaskArn           string             `json:"taskArn"`
	NetworkInterfaces []NetworkInterface `json:"networkInterfaces"`
	CPU               string             `json:"cpu"`
	ExitCode          int                `json:"exitCode"`
	Reason            string             `json:"reason"`
}

// ECSTaskStateChangeEvent is the event indicates ECS task state is changed.
type ECSTaskStateChangeEvent struct {
	Attachments       []TaskAttachment `json:"attachments"`
	AvailabilityZone  string           `json:"availabilityZone"`
	ClusterArn        string           `json:"clusterArn"`
	Containers        []Container      `json:"containers"`
	CreatedAt         time.Time        `json:"createdAt"`
	LaunchType        string           `json:"launchType"`
	CPU               string           `json:"cpu"`
	Memory            string           `json:"memory"`
	DesiredStatus     string           `json:"desiredStatus"`
	Group             string           `json:"group"`
	LastStatus        string           `json:"lastStatus"`
	Connectivity      string           `json:"connectivity"`
	ConnectivityAt    time.Time        `json:"connectvityAt"`
	PullStartedAt     time.Time        `json:"pullStartedAt"`
	StartedAt         time.Time        `json:"startedAt"`
	PullStoppedAt     time.Time        `json:"pullStoppedAt"`
	UpdatedAt         time.Time        `json:"updatedAt"`
	TaskArn           string           `json:"taskArn"`
	TaskDefinitionArn string           `json:"taskDefinitionArn"`
	Version           int              `json:"version"`
	PlatformVersion   string           `json:"platformVersion"`
	StoppedReason     string           `json:"stoppedReason"`
	StopCode          string           `json:"stopCode"`
	StoppingAt        time.Time        `json:"stoppingAt"`
	StoppedAt         time.Time        `json:"stoppedAt"`
}
