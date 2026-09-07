# downward

A small metadata contract for Docker Swarm applications, inspired by the Kubernetes Downward API.

**downward** defines a standard set of environment variables that expose information about the current Swarm stack, service, task, and node to an application running inside a container.

## Environment contract

| Environment variable | Description          | Docker Swarm source  |
| -------------------- | -------------------- | -------------------- |
| `SWARM_STACK_NAME`   | Stack name           | deployment context   |
| `SWARM_SERVICE_ID`   | Service ID           | `{{.Service.ID}}`    |
| `SWARM_SERVICE_NAME` | Service name         | `{{.Service.Name}}`  |
| `SWARM_TASK_ID`      | Task ID              | `{{.Task.ID}}`       |
| `SWARM_TASK_NAME`    | Task name            | `{{.Task.Name}}`     |
| `SWARM_TASK_SLOT`    | Replicated task slot | `{{.Task.Slot}}`     |
| `SWARM_NODE_ID`      | Swarm node ID        | `{{.Node.ID}}`       |
| `SWARM_NODE_NAME`    | Swarm node name      | `{{.Node.Hostname}}` |

`SWARM_TASK_SLOT` is optional in the SDK because a slot is not meaningful for every type of Swarm service.
