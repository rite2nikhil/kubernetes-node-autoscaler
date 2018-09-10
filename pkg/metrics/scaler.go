package metrics

import (
	"fmt"
	"sort"
)

type Scaler interface {
	Up(count int64, exclude []string) []string
	Down(count int64, exclude []string) []string
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

func (s *NodeScaler) Down(count int, exclude []string) ([]string, error) {
	if err := s.nodestore.Refresh(); err != nil {
		return nil, err
	}
	nodes := s.nodestore.GetAllNodes()
	if len(nodes) == 0 {
		return nil, fmt.Errorf("nodes not found")
	}
	nodes = excludeNodes(nodes, s.ScaleDown.ExcludedNodes)
	result := make([]string, count)
	for i := 0; i < count; i++ {
		allNodeMetrics := make(map[string]map[Metric]*RankResult)
		for _, m := range s.ScaleDown.Metrics {
			mt := &NodeCPU{nodestore: s.nodestore}
			results, err := RankByValue(nodes, mt)
			if err != nil {
				return nil, err
			}
			for node, result := range results {
				allNodeMetrics[node] = make(map[Metric]*RankResult)
				allNodeMetrics[node][m] = result
			}
		}
		scores := make([]RankResult, len(nodes))
		k := 0
		for n, mrs := range allNodeMetrics {
			for m, y := range mrs {
				scores[k].Key = n
				scores[k].RealValue = y.Rank * m.Weight
			}
			k++
		}
		sort.Sort(ByValue(scores))
		selected := scores[len(scores)-1].Key
		result = append(result, selected)
		nodes = excludeNodes(nodes, []string{selected})
	}
	return result, nil
}

func excludeNodes(a []string, exclude []string) (b []string) {
	em := make(map[string]bool)
	for _, e := range exclude {
		em[e] = true
	}
	b = a[:0]
	for _, x := range a {
		if !em[x] {
			b = append(b, x)
		}
	}
	return b
}
