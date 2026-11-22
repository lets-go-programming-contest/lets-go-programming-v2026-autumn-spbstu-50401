package main

import (
	"encoding/xml"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func LoadConfig(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func ParseXML(path string) []Valute {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var vc ValCurs
	if err := xml.Unmarshal(data, &vc); err != nil {
		panic(err)
	}

	return vc.Valutes
}

func parseValue(s string) float64 {
	s = strings.ReplaceAll(s, ",", ".")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func sortByValueDesc(valutes []Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return parseValue(valutes[i].Value) > parseValue(valutes[j].Value)
	})
}

func main() {
	configPath := flag.String("config", "", "path to yaml config")
	flag.Parse()

	if *configPath == "" {
		panic("config path is required (-config)")
	}

	cfg := LoadConfig(*configPath)
	valutes := ParseXML(cfg.InputFile)
	sortByValueDesc(valutes)

}
