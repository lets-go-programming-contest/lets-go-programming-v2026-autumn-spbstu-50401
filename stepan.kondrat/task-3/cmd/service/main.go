package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/net/html/charset"
)

const (
	dirPermissions  = 0o750
	filePermissions = 0o600
)

var ErrInvalidConfig = errors.New("configuration must contain non-empty input-file and output-file")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Currency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyRates struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

type CurrencyResult struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "Path to YAML configuration")

	flag.Parse()

	if *configPath == "" {
		panic("configuration path is required: use --config")
	}

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := loadCurrencies(config.InputFile)
	if err != nil {
		panic(err)
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	err = saveCurrencies(config.OutputFile, currencies)
	if err != nil {
		panic(err)
	}
}

func loadConfig(path string) (Config, error) {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("reading configuration: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("decoding configuration: %w", err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return config, ErrInvalidConfig
	}

	return config, nil
}

func loadCurrencies(path string) ([]CurrencyResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading XML file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var rates CurrencyRates

	err = decoder.Decode(&rates)
	if err != nil {
		return nil, fmt.Errorf("decoding XML: %w", err)
	}

	results := make([]CurrencyResult, 0, len(rates.Currencies))

	for _, currency := range rates.Currencies {
		valueText := strings.ReplaceAll(currency.Value, ",", ".")

		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil {
			return nil, fmt.Errorf(
				"parsing value for currency %q: %w",
				currency.CharCode,
				err,
			)
		}

		results = append(results, CurrencyResult{
			NumCode:  currency.NumCode,
			CharCode: currency.CharCode,
			Value:    value,
		})
	}

	return results, nil
}

func saveCurrencies(path string, currencies []CurrencyResult) error {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("encoding currencies as JSON: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), dirPermissions)
	if err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	err = os.WriteFile(path, data, filePermissions)
	if err != nil {
		return fmt.Errorf("writing JSON file: %w", err)
	}

	return nil
}
