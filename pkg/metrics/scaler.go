package metrics

type interface Scaler {
	Up(count int, exclude []string) []string 
	Down(count int, exclude []string) []string 
}

type NodeScaler struct {
	ScaleUp   ScaleConfig
	ScaleDown ScaleConfig
	nodestore NodeStore
}

func (s *NodeScaler) Up(count int) []string {
	return nil
}

func (s *NodeScaler) Down(count int, exclude []string) []string {
	nodes := filterNs.nodestore.GetAllNodes()
	nodes = excludeNodes(nodes, s.ScaleDown.ExcludedNodes)
	s.nodestore.Refresh()
	for i = 0; i < count; i++ {
		for m := range s.ScaleDown.Metrics {

		}
	}
}

func excludeNodes(a []string, exclude []string) (b []string) {
	em := make(map[string]bool)
	for _, e := range exclude {
		em[e] = true
	}
	b := a[:0]
	for _, x := range a {
		if em[x] {
			b = append(b, x)
		}
	}
}
