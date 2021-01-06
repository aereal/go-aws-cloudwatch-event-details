package cweventdetails

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"
)

func TestParseEventDetail(t *testing.T) {
	tests := []struct {
		name         string
		testDataPath string
		want         interface{}
		wantErr      bool
	}{
		{"ECS/task state change/ok", "testdata/ecs.ok.json", ecsTaskStateChangeOK, false},
		{"ECS/container instance/ok", "testdata/ecs.container-instance.ok.json", ecsContainerInstanceStateChangeOK, false},
		{"ECS/deployment/in progress", "testdata/ecs.deployment.in-progress.json", ecsDeploymentInProgress, false},
		{"ECS/deployment/rollback", "testdata/ecs.deployment.rollback.json", ecsDeploymentRollback, false},
		{"ECS/deployment/completed", "testdata/ecs.deployment.completed.json", ecsDeploymentCompleted, false},
		{"ECS/deployment/failed", "testdata/ecs.deployment.failed.json", ecsDeploymentFailed, false},
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testDataPath == "" {
				t.Fatal("testDataPath is empty")
			}
			f, err := os.Open(filepath.Join(cwd, tt.testDataPath))
			if err != nil {
				t.Fatal(err)
			}
			var ev events.CloudWatchEvent
			if err := json.NewDecoder(f).Decode(&ev); err != nil {
				t.Fatal(err)
			}
			got, err := ParseEventDetail(ev)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEventDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ParseEventDetail(): (-want +got):\n%s", diff)
			}
		})
	}
}

func mustParseTime(repr string) time.Time {
	parsed, err := time.Parse(time.RFC3339Nano, repr)
	if err != nil {
		panic(err)
	}
	return parsed
}

var (
	ecsTaskStateChangeOK = &ECSTaskStateChangeEvent{
		Attachments: []TaskAttachment{
			{
				ID:     "1789bcae-ddfb-4d10-8ebe-8ac87ddba5b8",
				Type:   "eni",
				Status: "ATTACHED",
				Details: []TaskAttachmentDetail{
					{
						Name:  "subnetId",
						Value: "subnet-abcd1234",
					},
					{
						Name:  "networkInterfaceId",
						Value: "eni-abcd1234",
					},
					{
						Name:  "macAddress",
						Value: "0a:98:eb:a7:29:ba",
					},
					{
						Name:  "privateIPv4Address",
						Value: "10.0.0.139",
					},
				},
			},
		},
		AvailabilityZone: "us-west-2c",
		ClusterArn:       "arn:aws:ecs:us-west-2:111122223333:cluster/FargateCluster",
		Containers: []Container{
			{
				ContainerArn: "arn:aws:ecs:us-west-2:111122223333:container/cf159fd6-3e3f-4a9e-84f9-66cbe726af01",
				LastStatus:   "RUNNING",
				Name:         "FargateApp",
				Image:        "111122223333.dkr.ecr.us-west-2.amazonaws.com/hello-repository:latest",
				ImageDigest:  "sha256:74b2c688c700ec95a93e478cdb959737c148df3fbf5ea706abe0318726e885e6",
				RuntimeID:    "ad64cbc71c7fb31c55507ec24c9f77947132b03d48d9961115cf24f3b7307e1e",
				TaskArn:      "arn:aws:ecs:us-west-2:111122223333:task/FargateCluster/c13b4cb40f1f4fe4a2971f76ae5a47ad",
				NetworkInterfaces: []NetworkInterface{
					{
						AttachmentID:       "1789bcae-ddfb-4d10-8ebe-8ac87ddba5b8",
						PrivateIpv4Address: "10.0.0.139",
					},
				},
				CPU: "0",
			},
		},
		LaunchType:        "FARGATE",
		CPU:               "256",
		Memory:            "512",
		DesiredStatus:     "RUNNING",
		Group:             "family:sample-fargate",
		LastStatus:        "RUNNING",
		Connectivity:      "CONNECTED",
		TaskArn:           "arn:aws:ecs:us-west-2:111122223333:task/FargateCluster/c13b4cb40f1f4fe4a2971f76ae5a47ad",
		TaskDefinitionArn: "arn:aws:ecs:us-west-2:111122223333:task-definition/sample-fargate:1",
		Version:           4,
		PlatformVersion:   "1.3.0",
		ConnectivityAt:    mustParseTime("2020-01-23T17:57:38.453Z"),
		PullStartedAt:     mustParseTime("2020-01-23T17:57:52.103Z"),
		PullStoppedAt:     mustParseTime("2020-01-23T17:57:55.103Z"),
		StartedAt:         mustParseTime("2020-01-23T17:57:58.103Z"),
		CreatedAt:         mustParseTime("2020-01-23T17:57:34.402Z"),
		UpdatedAt:         mustParseTime("2020-01-23T17:57:58.103Z"),
	}
	ecsContainerInstanceStateChangeOK = &ECSContainerInstanceStateChangeEvent{
		AgentConnected: true,
		Attributes: []ECSContainerInstanceAttribute{
			{Name: "com.amazonaws.ecs.capability.logging-driver.syslog"},
			{Name: "com.amazonaws.ecs.capability.task-iam-role-network-host"},
			{Name: "com.amazonaws.ecs.capability.logging-driver.awslogs"},
			{Name: "com.amazonaws.ecs.capability.logging-driver.json-file"},
			{Name: "com.amazonaws.ecs.capability.docker-remote-api.1.17"},
			{Name: "com.amazonaws.ecs.capability.privileged-container"},
			{Name: "com.amazonaws.ecs.capability.docker-remote-api.1.18"},
			{Name: "com.amazonaws.ecs.capability.docker-remote-api.1.19"},
			{Name: "com.amazonaws.ecs.capability.ecr-auth"},
			{Name: "com.amazonaws.ecs.capability.docker-remote-api.1.20"},
			{Name: "com.amazonaws.ecs.capability.docker-remote-api.1.21"},
			{Name: "com.amazonaws.ecs.capability.docker-remote-api.1.22"},
			{Name: "com.amazonaws.ecs.capability.docker-remote-api.1.23"},
			{Name: "com.amazonaws.ecs.capability.task-iam-role"},
		},
		ClusterArn:           "arn:aws:ecs:us-east-1:111122223333:cluster/default",
		ContainerInstanceArn: "arn:aws:ecs:us-east-1:111122223333:container-instance/b54a2a04-046f-4331-9d74-3f6d7f6ca315",
		EC2InstanceID:        "i-f3a8506b",
		RegisteredResources: []ECSResource{
			{Name: "CPU", Type: "INTEGER", IntegerValue: 2048},
			{Name: "MEMORY", Type: "INTEGER", IntegerValue: 3767},
			{Name: "PORTS", Type: "STRINGSET", StringSetValue: []string{"22", "2376", "2375", "51678", "51679"}},
			{Name: "PORTS_UDP", Type: "STRINGSET", StringSetValue: []string{}},
		},
		RemainingResources: []ECSResource{
			{Name: "CPU", Type: "INTEGER", IntegerValue: 1988},
			{Name: "MEMORY", Type: "INTEGER", IntegerValue: 767},
			{Name: "PORTS", Type: "STRINGSET", StringSetValue: []string{"22", "2376", "2375", "51678", "51679"}},
			{Name: "PORTS_UDP", Type: "STRINGSET", StringSetValue: []string{}},
		},
		Status:  "ACTIVE",
		Version: 14801,
		VersionInfo: ECSContainerAgentInfo{
			AgentHash:     "aebcbca",
			AgentVersion:  "1.13.0",
			DockerVersion: "DockerVersion: 1.11.2",
		},
		UpdatedAt: mustParseTime("2016-12-06T16:41:06.991Z"),
	}
	ecsDeploymentInProgress = &ECSDeploymentStateChangeEvent{
		EventType:    "INFO",
		EventName:    "SERVICE_DEPLOYMENT_IN_PROGRESS",
		DeploymentID: "ecs-svc/123",
		UpdatedAt:    mustParseTime("2020-05-23T11:11:11Z"),
		Reason:       "ECS deployment deploymentId in progress.",
	}
	ecsDeploymentRollback = &ECSDeploymentStateChangeEvent{
		EventType:    "INFO",
		EventName:    "SERVICE_DEPLOYMENT_IN_PROGRESS",
		DeploymentID: "ecs-svc/123",
		UpdatedAt:    mustParseTime("2020-05-23T11:11:11Z"),
		Reason:       "ECS deployment circuit breaker: rolling back to deploymentId deploymentID.",
	}
	ecsDeploymentCompleted = &ECSDeploymentStateChangeEvent{
		EventType:    "INFO",
		EventName:    "SERVICE_DEPLOYMENT_COMPLETED",
		DeploymentID: "ecs-svc/123",
		UpdatedAt:    mustParseTime("2020-05-23T11:11:11Z"),
		Reason:       "ECS deployment deploymentID completed.",
	}
	ecsDeploymentFailed = &ECSDeploymentStateChangeEvent{
		EventType:    "ERROR",
		EventName:    "SERVICE_DEPLOYMENT_FAILED",
		DeploymentID: "ecs-svc/123",
		UpdatedAt:    mustParseTime("2020-05-23T11:11:11Z"),
		Reason:       "ECS deployment circuit breaker: task failed to start.",
	}
)
