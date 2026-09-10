# downward

A small metadata contract for Docker Swarm applications, inspired by the Kubernetes Downward API.

**downward** defines a standard set of environment variables that expose information about the current Swarm stack, service, task, and node to an application running inside a container.

## Environment contract

| Environment variable | Description          | Docker Swarm source  |
| -------------------- | -------------------- | -------------------- |
| `SWARM_STACK_NAME`   | Stack name           | `{{index .Service.Labels "com.docker.stack.namespace"}}`    |
| `SWARM_SERVICE_ID`   | Service ID           | `{{.Service.ID}}`    |
| `SWARM_SERVICE_NAME` | Service name         | `{{.Service.Name}}`  |
| `SWARM_TASK_ID`      | Task ID              | `{{.Task.ID}}`       |
| `SWARM_TASK_NAME`    | Task name            | `{{.Task.Name}}`     |
| `SWARM_TASK_SLOT`    | Replicated task slot | `{{.Task.Slot}}`     |
| `SWARM_NODE_ID`      | Swarm node ID        | `{{.Node.ID}}`       |
| `SWARM_NODE_NAME`    | Swarm node name      | `{{.Node.Hostname}}` |

`SWARM_TASK_SLOT` is optional in the SDK because a slot is not meaningful for every type of Swarm service.

When deploying with [swarm-deploy](https://github.com/swarm-deploy/swarm-deploy), these variables can be injected automatically. swarm-deploy adds the environment contract to services, while Docker Swarm resolves the service, task, and node placeholders at runtime.

## Go

Install the package:

```bash
go get github.com/swarm-deploy/downward/go
```

Load metadata from the environment:

```go
package main

import (
	"fmt"
	"log"

	"github.com/swarm-deploy/downward/go"
)

func main() {
	info, err := downward.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("stack: %s\n", info.Stack.Name)
	fmt.Printf("service: %s (%s)\n", info.Service.Name, info.Service.ID)
	fmt.Printf("task: %s (%s)\n", info.Task.Name, info.Task.ID)

	if info.Task.Slot != nil {
		fmt.Printf("task slot: %d\n", *info.Task.Slot)
	}

	fmt.Printf("node: %s (%s)\n", info.Node.Name, info.Node.ID)
}
```

## OpenTelemetry

Use [downward-otel](https://github.com/swarm-deploy/downward-otel) to expose downward metadata as OpenTelemetry resource attributes.
