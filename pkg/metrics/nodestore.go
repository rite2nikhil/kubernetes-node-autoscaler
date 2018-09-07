package metrics

import (
	"fmt"

	"k8s.io/api/core/v1"
)

type NodeStore interface {
	Refresh() (err error)
	Get(node string) (obj *v1.Node, err error)
	// GetByLabel(label string) (obj map[Label][]NodeName, err error)
	GetAllNodes() (nodes []string)
}

type NodeCache struct {
	nodecache map[string]*v1.Node
	k8s       KubeClient
}

func (c *NodeCache) Refresh() (err error) {
	if node == nil {
		return fmt.Errorf("node name is nil")
	}
	objs, err := k8s.KubeClient.ClientSet.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	if len(objs) == 0 {
		return fmt.Errorf("nodes not found")
	}
	for o := range objs {
		nodecache[o.Spec.name] = o
	}

	return nil
}

func (c *NodeCache) GetAllNodes() (nodes []string) {
	for k, _ := range c.nodecache {
		nodes = append(nodes, k)
	}

	return nodes
}

func (c *NodeCache) Get(node string) (obj *v1.Node, err error) {
	if node == nil {
		return nil, fmt.Errorf("node name is nil")
	}
	if obj, ok := nodecache[node]; !ok {
		return nil, fmt.Errorf("node not found, node: %", node)
	}

	return obj, nil
}
