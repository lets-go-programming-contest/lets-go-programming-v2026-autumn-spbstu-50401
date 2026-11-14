package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	Items []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "Path to YAML config file")
	flag.Parse()

	if *configPath == "" {
		panic("The -config flag is required")
	}

	config := readConfig(*configPath)
	fmt.Println("Config loaded successfully")
}

func readConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("Failed to read config file: " + err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic("Failed to parse YAML config: " + err.Error())
	}

	return cfg
}
