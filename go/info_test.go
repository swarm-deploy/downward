package downward

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Setenv(EnvStackName, "stack")
	t.Setenv(EnvServiceID, "service-id")
	t.Setenv(EnvServiceName, "service")
	t.Setenv(EnvTaskID, "task-id")
	t.Setenv(EnvTaskName, "task")
	t.Setenv(EnvTaskSlot, "3")
	t.Setenv(EnvNodeID, "node-id")
	t.Setenv(EnvNodeName, "node")

	info, err := Load()
	require.NoError(t, err)
	require.Equal(t, "stack", info.Stack.Name)
	require.Equal(t, "service-id", info.Service.ID)
	require.Equal(t, "service", info.Service.Name)
	require.Equal(t, "task-id", info.Task.ID)
	require.Equal(t, "task", info.Task.Name)
	require.NotNil(t, info.Task.Slot)
	require.Equal(t, 3, *info.Task.Slot)
	require.Equal(t, "node-id", info.Node.ID)
	require.Equal(t, "node", info.Node.Name)
}

func TestLoadWithoutTaskSlot(t *testing.T) {
	info, err := Load()
	require.NoError(t, err)
	require.Nil(t, info.Task.Slot)
}

func TestLoadInvalidTaskSlot(t *testing.T) {
	t.Setenv(EnvTaskSlot, "invalid")

	_, err := Load()
	require.Error(t, err)
}
