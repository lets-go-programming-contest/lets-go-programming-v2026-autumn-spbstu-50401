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

func readConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		panic("invalid configuration")
	}

	return config
}

func main() {
	configPath := flag.String("config", "", "configuration file path")
	flag.Parse()

	if *configPath == "" {
		panic("config path is required")
	}

	config := readConfig(*configPath)
	_ = config
}
