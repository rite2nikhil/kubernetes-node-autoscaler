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
	down := metrics.ScaleConfig{}
	scaler := metrics.NewNodeScaler(k8s, up, down)
	scaler.Down(1, nil)
}
