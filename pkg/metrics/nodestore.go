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
	k8s       KubeNodeClient
}

func NewNodeCache(k8s KubeNodeClient) (nc *NodeCache) {
	return &NodeCache{k8s: k8s, nodecache: make(map[string]*v1.Node)}
}

func (c *NodeCache) Refresh() (err error) {
	objs, err := c.k8s.ListNodes()
	if err != nil {
		return err
	}
	if len(objs.Items) == 0 {
		return fmt.Errorf("nodes not found")
	}
	for _, o := range objs.Items {
		c.nodecache[o.GetName()] = &o
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
	if obj, ok := c.nodecache[node]; ok {
		return obj, nil
	}
	return nil, fmt.Errorf("node not found, node: %s", node)
}
