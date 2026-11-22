package main

import (
	"encoding/xml"
	"flag"
	"os"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
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

func loadXML(path string) ValCurs {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(err)
	}

	return valCurs
}

func main() {
	configPath := flag.String("config", "", "path to yaml config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is required")
	}

	cfg := loadConfig(*configPath)
	_ = loadXML(cfg.InputFile)
}
