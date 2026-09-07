package downward

import (
	"fmt"
	"os"
	"strconv"
)

type Info struct {
	Stack   Stack   `json:"stack"`
	Service Service `json:"service"`
	Task    Task    `json:"task"`
	Node    Node    `json:"node"`
}

func Load() (Info, error) {
	var slot *int
	if value, ok := os.LookupEnv(EnvTaskSlot); ok {
		parsedSlot, err := strconv.Atoi(value)
		if err != nil {
			return Info{}, fmt.Errorf("parse %s: %w", EnvTaskSlot, err)
		}
		slot = &parsedSlot
	}

	return Info{
		Stack: Stack{
			Name: os.Getenv(EnvStackName),
		},
		Service: Service{
			ID:   os.Getenv(EnvServiceID),
			Name: os.Getenv(EnvServiceName),
		},
		Task: Task{
			ID:   os.Getenv(EnvTaskID),
			Name: os.Getenv(EnvTaskName),
			Slot: slot,
		},
		Node: Node{
			ID:   os.Getenv(EnvNodeID),
			Name: os.Getenv(EnvNodeName),
		},
	}, nil
}

func (i Info) ToMap() map[string]any {
	return map[string]any{
		EnvStackName:   i.Stack.Name,
		EnvServiceID:   i.Service.ID,
		EnvServiceName: i.Service.Name,
		EnvTaskID:      i.Task.ID,
		EnvTaskName:    i.Task.Name,
		EnvTaskSlot:    i.Task.Slot,
		EnvNodeID:      i.Node.ID,
		EnvNodeName:    i.Node.Name,
	}
}

type Stack struct {
	Name string `env:"SWARM_STACK_NAME" json:"name"`
}

func (s Stack) ToMap() map[string]any {
	return map[string]any{
		EnvStackName: s.Name,
	}
}

type Service struct {
	ID   string `env:"SWARM_SERVICE_ID" json:"id"`
	Name string `env:"SWARM_SERVICE_NAME" json:"name"`
}

func (s Service) ToMap() map[string]any {
	return map[string]any{
		EnvServiceID:   s.ID,
		EnvServiceName: s.Name,
	}
}

type Task struct {
	ID   string `env:"SWARM_TASK_ID" json:"id"`
	Name string `env:"SWARM_TASK_NAME" json:"name"`
	Slot *int   `env:"SWARM_TASK_SLOT" json:"slot"`
}

func (t Task) ToMap() map[string]any {
	return map[string]any{
		EnvTaskID:   t.ID,
		EnvTaskName: t.Name,
		EnvTaskSlot: t.Slot,
	}
}

type Node struct {
	ID   string `env:"SWARM_NODE_ID" json:"id"`
	Name string `env:"SWARM_NODE_NAME" json:"name"`
}

func (n Node) ToMap() map[string]any {
	return map[string]any{
		EnvNodeID:   n.ID,
		EnvNodeName: n.Name,
	}
}
