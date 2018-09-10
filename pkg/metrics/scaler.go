package metrics

import (
	"fmt"
)

type Scaler interface {
	Up(count int, exclude []string) []string
	Down(count int, exclude []string) []string
}

type NodeScaler struct {
	ScaleUp   ScaleConfig
	ScaleDown ScaleConfig
	nodestore NodeStore
}

func NewNodeScaler(k8s KubeNodeClient, up, down ScaleConfig) *NodeScaler {
	ns := NewNodeCache(k8s)
	return &NodeScaler{nodestore: ns, ScaleUp: up, ScaleDown: down}
}

func (s *NodeScaler) Up(count int) []string {
	return nil
}

func (s *NodeScaler) Down(count int, exclude []string) []string {
	nodes := s.nodestore.GetAllNodes()
	nodes = excludeNodes(nodes, s.ScaleDown.ExcludedNodes)
	result := make([]string, count)
	s.nodestore.Refresh()
	for i := 0; i < count; i++ {
		for _, m := range s.ScaleDown.Metrics {
			fmt.Printf("%v", m)
			mt := &NodeCPU{nodestore: s.nodestore}
			r, _ := RankByValue(nodes, mt)
			fmt.Printf("%v", r)
		}
	}

	return result
}

func excludeNodes(a []string, exclude []string) (b []string) {
	em := make(map[string]bool)
	for _, e := range exclude {
		em[e] = true
	}
	b = a[:0]
	for _, x := range a {
		if em[x] {
			b = append(b, x)
		}
	}
	return b
}
