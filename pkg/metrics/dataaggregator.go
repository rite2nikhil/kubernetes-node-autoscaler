package metrics

type DataAggregator interface {
	Get(key string) interface{}, error
}

type NodeCPU struct {
	nodestore NodeStore
}

// GetCurrentCPU get node current cpu for from apiserver
func (c *NodeCPU) Get(key string) int64, error {
	n, err := nodestore.GetNode(node)
	if err != nil {
		return -1, err
	}
	if n.Status == nil || n.Status.Allocatable == nil {
		return -1, fmt.Errorf("node is nil, node: %", node)
	}
	cpu, ok = n.Status.Allocatable[v1.ResourceCPU]
	if !ok {
		return -1, fmt.Errorf("cpu resource not found, node %s", node)
	}
	return cpu.Value(), nil
}