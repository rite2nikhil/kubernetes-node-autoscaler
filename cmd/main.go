package main

import (
	"fmt"

	metrics "github.com/rite2nikhil/kubernetes-node-scaler/pkg/metrics"
)

func main() {
	k8s, err := metrics.NewKubeClient()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("starting")
	up := metrics.ScaleConfig{}
	m := []metrics.Metric{
		{Name: "cpu", Weight: 1},
	}
	down := metrics.ScaleConfig{m, []string{}}

	scaler := metrics.NewNodeScaler(k8s, up, down)
	x, e := scaler.Down(10, nil)
	fmt.Printf("result: %v %v", x, e)
}
