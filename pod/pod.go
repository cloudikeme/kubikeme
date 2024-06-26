package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	pendinG State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Pod struct {
	ID            uuid.UUID
	Name          string
	state         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBinding   map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

type PodEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Pod       Pod
}
