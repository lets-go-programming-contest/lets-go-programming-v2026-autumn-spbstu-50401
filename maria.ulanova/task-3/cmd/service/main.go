package main

import (
	"encoding/xml"
	"flag"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
	"os"
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

type Currency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	Currencies []Currency `xml:"Valute"`
}

func readXML(path string) ValCurs {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data ValCurs
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}

	return data
}

func main() {
	configPath := flag.String("config", "", "configuration file path")
	flag.Parse()

	if *configPath == "" {
		panic("config path is required")
	}

	config := readConfig(*configPath)
	currencies := readXML(config.InputFile)
	_ = currencies
}
