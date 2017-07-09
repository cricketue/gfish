package config

import (
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Addr string `yaml:"addr"`
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
