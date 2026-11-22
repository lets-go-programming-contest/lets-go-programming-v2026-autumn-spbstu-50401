package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func loadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("input-file and output-file must be specified in config")
	}

	return cfg
}

func main() {
	configPath := flag.String("config", "", "path to yaml config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is required")
	}

	_ = loadConfig(*configPath)
}
