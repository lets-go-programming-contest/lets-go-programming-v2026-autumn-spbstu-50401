package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v4"
	"golang.org/x/net/html/charset"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o644
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type CommaFloat float64

func (f *CommaFloat) UnmarshalText(text []byte) error {
	normalized := strings.Replace(string(text), ",", ".", 1)

	value, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("failed to parse number %q: %w", text, err)
	}

	*f = CommaFloat(value)

	return nil
}

type Currency struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

type ExchangeRates struct {
	Currencies []Currency `xml:"Valute"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &cfg, nil
}

func loadCurrencies(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var rates ExchangeRates
	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("failed to decode input file: %w", err)
	}

	return rates.Currencies, nil
}

func sortByValue(currencies []Currency) {
	slices.SortFunc(currencies, func(a, b Currency) int {
		return cmp.Compare(b.Value, a.Value)
	})
}

func saveCurrencies(path string, currencies []Currency) error {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to encode currencies: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), dirPermissions); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := loadCurrencies(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sortByValue(currencies)

	if err := saveCurrencies(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
