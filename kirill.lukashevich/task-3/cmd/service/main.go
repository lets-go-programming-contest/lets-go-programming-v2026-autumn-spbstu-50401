package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrInvalidConfig = errors.New("input-file and output-file must not be empty")

func main() {
	configPath := flag.String("config", "config.yaml", "Path to YAML config")

	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := readCurrencies(config.InputFile)
	if err != nil {
		panic(err)
	}

	_ = currencies
}

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func loadConfig(path string) (Config, error) {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read config: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("decode config: %w", err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return config, ErrInvalidConfig
	}

	return config, nil
}

type Currency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyList struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Items   []Currency `xml:"Valute"`
}

func readCurrencies(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))

	var currencies CurrencyList

	err = decoder.Decode(&currencies)
	if err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	return currencies.Items, nil
}
