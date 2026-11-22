package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
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

type OutputCurrency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
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

func SaveJSON(path string, valutes []Valute) {
	out := make([]OutputCurrency, 0, len(valutes))
	for _, v := range valutes {
		num, err := strconv.Atoi(v.NumCode)
		if err != nil {
			panic(err)
		}
		out = append(out, OutputCurrency{
			NumCode:  num,
			CharCode: v.CharCode,
			Value:    parseValue(v.Value),
		})
	}

	data, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		panic(err)
	}

	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			panic(err)
		}
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		panic(err)
	}
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
	SaveJSON(cfg.OutputFile, valutes)

}
