package metrics

import (
	"encoding/json"
	"fmt"
	"os"
)

type ScaleConfig struct {
	Metrics []struct {
		Name   string `json:"name"`
		Weight string `json:"weight"`
	} `json:"metrics"`
	ExcludedNodes []string `json:"excludednodes"`
}

func LoadConfiguration(file string) Config {
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
