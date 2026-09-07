package downward

type Info struct {
	Stack   Stack
	Service Service
	Task    Task
	Node    Node
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
	Name string `env:"SWARM_STACK_NAME"`
}

func (s Stack) ToMap() map[string]any {
	return map[string]any{
		EnvStackName: s.Name,
	}
}

type Service struct {
	ID   string `env:"SWARM_SERVICE_ID"`
	Name string `env:"SWARM_SERVICE_NAME"`
}

func (s Service) ToMap() map[string]any {
	return map[string]any{
		EnvServiceID:   s.ID,
		EnvServiceName: s.Name,
	}
}

type Task struct {
	ID   string `env:"SWARM_TASK_ID"`
	Name string `env:"SWARM_TASK_NAME"`
	Slot *int   `env:"SWARM_TASK_SLOT"`
}

func (t Task) ToMap() map[string]any {
	return map[string]any{
		EnvTaskID:   t.ID,
		EnvTaskName: t.Name,
		EnvTaskSlot: t.Slot,
	}
}

type Node struct {
	ID   string `env:"SWARM_NODE_ID"`
	Name string `env:"SWARM_NODE_NAME"`
}

func (n Node) ToMap() map[string]any {
	return map[string]any{
		EnvNodeID:   n.ID,
		EnvNodeName: n.Name,
	}
}
