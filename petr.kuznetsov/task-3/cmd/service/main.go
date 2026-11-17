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
	dirPermissions  = 0o755
	filePermissions = 0o644
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

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("input-file and output-file must be set")
	}

	return cfg
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func loadCurrencies(path string) []Currency {
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

	result := make([]Currency, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		num := 0
		numCode := strings.TrimSpace(valute.NumCode)

		if numCode != "" {
			num, err = strconv.Atoi(numCode)
			if err != nil {
				panic(err)
			}
		}

		val, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(valute.Value), ",", ".", 1), 64)
		if err != nil {
			panic(err)
		}

		result = append(result, Currency{num, strings.TrimSpace(valute.CharCode), val})
	}

	return result
}

func saveJSON(path string, data []Currency) {
	if err := os.MkdirAll(filepath.Dir(path), dirPermissions); err != nil {
		panic(err)
	}

	serialized, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, serialized, filePermissions); err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "", "path to the YAML config")
	flag.Parse()

	if *configPath == "" {
		panic("Doesn't add flag -config")
	}

	cfg := loadConfig(*configPath)
	currencies := loadCurrencies(cfg.InputFile)

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	saveJSON(cfg.OutputFile, currencies)
}
