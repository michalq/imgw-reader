package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type CommonConfig struct {
	Db DbConfig `json:"db"`
}

type DbConfig struct {
	Url string `json:"url"`
}

func Read(filename string) (*CommonConfig, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file %s error: %v", filename, err)
	}
	var c CommonConfig
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
	}
	return &c, nil
}
