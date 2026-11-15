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

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

const (
	dirPermissions  = 0o750
	filePermissions = 0o600
)

var (
	ErrConfigPathRequired = errors.New("use (--config or -config) to indicate configuration path")
	ErrInputFileRequired  = errors.New("input-file is required")
	ErrOutputFileRequired = errors.New("output-file is required")
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

type OutputValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := loadXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	output, err := convertToOutput(valCurs.Valutes)
	if err != nil {
		panic(err)
	}

	sortOutput(output)

	if err := saveJSON(config.OutputFile, output); err != nil {
		panic(err)
	}
}

func loadConfig(path string) (Config, error) {
	var config Config

	if path == "" {
		return config, ErrConfigPathRequired
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read config %q: %w", path, err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("parse config %q: %w", path, err)
	}

	if config.InputFile == "" {
		return config, fmt.Errorf("config %q: %w", path, ErrInputFileRequired)
	}

	if config.OutputFile == "" {
		return config, fmt.Errorf("config %q: %w", path, ErrOutputFileRequired)
	}

	return config, nil
}

func loadXML(path string) (ValCurs, error) {
	var valCurs ValCurs

	data, err := os.ReadFile(path)
	if err != nil {
		return valCurs, fmt.Errorf("read input-file %q: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&valCurs); err != nil {
		return valCurs, fmt.Errorf("parse XML %q: %w", path, err)
	}

	return valCurs, nil
}

func parseValue(value string) (float64, error) {
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, ",", ".")

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("parse %q: %w", value, err)
	}

	return parsed, nil
}

func convertToOutput(valutes []Valute) ([]OutputValute, error) {
	output := make([]OutputValute, 0, len(valutes))

	for _, valute := range valutes {
		value, err := parseValue(valute.Value)
		if err != nil {
			return nil, fmt.Errorf("value of %q: %w", valute.CharCode, err)
		}

		output = append(output, OutputValute{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	return output, nil
}

func sortOutput(output []OutputValute) {
	sort.Slice(output, func(i, j int) bool {
		return output[i].Value > output[j].Value
	})
}

func saveJSON(path string, output []OutputValute) error {
	data, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal JSON: %w", err)
	}

	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, dirPermissions); err != nil {
			return fmt.Errorf("create dir %q: %w", dir, err)
		}
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write output-file %q: %w", path, err)
	}

	return nil
}
