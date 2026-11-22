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
	dirPerm  = 0o750
	filePerm = 0o600
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

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func loadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("Invalid config: input-file and output-file are required")
	}

	return cfg
}

func loadXML(path string) ValCurs {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}

	return valCurs
}

func convertAndSort(valCurs ValCurs) []Currency {
	currencies := make([]Currency, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		valueStr := strings.ReplaceAll(valute.Value, ",", ".")

		parsedValue, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			panic(err)
		}

		var currency Currency

		currency.NumCode = valute.NumCode
		currency.CharCode = valute.CharCode
		currency.Value = parsedValue

		currencies = append(currencies, currency)
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies
}

func saveJSON(path string, currencies []Currency) {
	outputDir := filepath.Dir(path)

	err := os.MkdirAll(outputDir, dirPerm)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path, jsonData, filePerm)
	if err != nil {
		panic(err)
	}
}

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to yaml config file")
	flag.Parse()

	if configPath == "" {
		panic("Config path is required")
	}

	cfg := loadConfig(configPath)
	valCurs := loadXML(cfg.InputFile)
	currencies := convertAndSort(valCurs)
	saveJSON(cfg.OutputFile, currencies)
}
