package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrInvalidConfig = errors.New("input-file and output-file must not be empty")

func main() {
	configPath := flag.String("config", "config.yaml", "Path to YAML config")

	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	_ = config
}

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func loadConfig(path string) (Config, error) {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read config: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("decode config: %w", err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return config, ErrInvalidConfig
	}

	return config, nil
}
