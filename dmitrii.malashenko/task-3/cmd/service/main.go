package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

const (
	dirPerm  = 0o755
	filePerm = 0o600
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
	NumCode  int    `xml:"NumCode"`
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

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	if err := decoder.Decode(&valCurs); err != nil {
		panic(err)
	}

	return valCurs.Valutes
}

func parseValue(s string) float64 {
	s = strings.ReplaceAll(strings.TrimSpace(s), ",", ".")

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return f
}

func toOutputCurrency(valute Valute) OutputCurrency {
	return OutputCurrency{
		NumCode:  valute.NumCode,
		CharCode: strings.TrimSpace(valute.CharCode),
		Value:    parseValue(valute.Value),
	}
}

func sortByValueDesc(valutes []Valute) {
	sort.SliceStable(valutes, func(i, j int) bool {
		return parseValue(valutes[i].Value) > parseValue(valutes[j].Value)
	})
}

func SaveJSON(path string, valutes []Valute) {
	out := make([]OutputCurrency, 0, len(valutes))

	for _, valute := range valutes {
		out = append(out, toOutputCurrency(valute))
	}

	data, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		panic(err)
	}

	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, dirPerm); err != nil {
			panic(err)
		}
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
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
