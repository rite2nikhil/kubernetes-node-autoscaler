package metrics

import (
	"fmt"

	"k8s.io/api/core/v1"
)

type DataAggregator interface {
	Get(key string) (interface{}, error)
}

type NodeCPU struct {
	nodestore NodeStore
}

// GetCurrentCPU get node current cpu for from apiserver
func (c *NodeCPU) Get(key string) (interface{}, error) {
	n, err := c.nodestore.Get(key)
	if err != nil {
		return -1, err
	}
	if n.Status.Allocatable == nil {
		return -1, fmt.Errorf("node is nil, node: %", key)
	}
	cpu, ok := n.Status.Allocatable[v1.ResourceCPU]
	if !ok {
		return -1, fmt.Errorf("cpu resource not found, node %s", key)
	}
	return cpu.Value(), nil
}
