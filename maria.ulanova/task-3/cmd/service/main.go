package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
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

type ResultCurrency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func sortCurrencies(data ValCurs) []ResultCurrency {
	result := make([]ResultCurrency, 0, len(data.Currencies))

	for _, currency := range data.Currencies {
		value := strings.ReplaceAll(currency.Value, ",", ".")
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(err)
		}

		result = append(result, ResultCurrency{
			NumCode:  currency.NumCode,
			CharCode: currency.CharCode,
			Value:    number,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

func saveJSON(path string, currencies []ResultCurrency) {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "", "configuration file path")
	flag.Parse()

	if *configPath == "" {
		panic("config path is required")
	}

	config := readConfig(*configPath)
	currencies := readXML(config.InputFile)
	result := sortCurrencies(currencies)
	saveJSON(config.OutputFile, result)
}
