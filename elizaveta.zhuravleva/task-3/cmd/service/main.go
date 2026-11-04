package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

const (
	directoryPermissions = 0o755
	filePermissions      = 0o600
)

var (
	errConfigPathRequired = errors.New("configuration path is required: use -config")
	errInputFileRequired  = errors.New("input-file is required")
	errOutputFileRequired = errors.New("output-file is required")
	errUnsupportedCharset = errors.New("unsupported XML encoding")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type XMLCurrencies struct {
	XMLName xml.Name      `xml:"ValCurs"`
	Items   []XMLCurrency `xml:"Valute"`
}

type XMLCurrency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "path to YAML configuration")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}

func run(configPath string) error {
	config, err := readConfig(configPath)
	if err != nil {
		return err
	}

	currencies, err := readCurrencies(config.InputFile)
	if err != nil {
		return err
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return writeCurrencies(config.OutputFile, currencies)
}

func readConfig(path string) (Config, error) {
	var config Config

	if strings.TrimSpace(path) == "" {
		return config, errConfigPathRequired
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read configuration: %w", err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("decode configuration: %w", err)
	}

	if strings.TrimSpace(config.InputFile) == "" {
		return config, errInputFileRequired
	}

	if strings.TrimSpace(config.OutputFile) == "" {
		return config, errOutputFileRequired
	}

	return config, nil
}

func readCurrencies(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charsetReader

	var source XMLCurrencies
	if err := decoder.Decode(&source); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	currencies := make([]Currency, 0, len(source.Items))

	for _, item := range source.Items {
		currency, err := convertCurrency(item)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(strings.TrimSpace(charset)) {
	case "windows-1251", "cp1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedCharset, charset)
	}
}

func convertCurrency(source XMLCurrency) (Currency, error) {
	var currency Currency

	numCodeText := strings.TrimSpace(source.NumCode)
	if numCodeText != "" {
		numCode, err := strconv.Atoi(numCodeText)
		if err != nil {
			return currency, fmt.Errorf("invalid currency code %q: %w", source.NumCode, err)
		}

		currency.NumCode = numCode
	}

	currency.CharCode = strings.TrimSpace(source.CharCode)

	valueText := strings.ReplaceAll(strings.TrimSpace(source.Value), ",", ".")

	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return currency, fmt.Errorf("invalid value for %s: %w", currency.CharCode, err)
	}

	currency.Value = value

	return currency, nil
}

func writeCurrencies(path string, currencies []Currency) error {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), directoryPermissions); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write JSON: %w", err)
	}

	return nil
}
