package metrics

import (
	"encoding/json"
	"fmt"
	"os"
)

type Metric struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

type ScaleConfig struct {
	Metrics       []Metric `json:"metrics"`
	ExcludedNodes []string `json:"excludednodes"`
}

func LoadConfiguration(file string) ScaleConfig {
	var config ScaleConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
