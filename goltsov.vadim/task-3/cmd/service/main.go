package main

import (
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
	outputDirPerm  = 0o750
	outputFilePerm = 0o600
)

var (
	ErrReadConfig       = errors.New("error reading config file")
	ErrParseConfig      = errors.New("error parsing config file")
	ErrOpenInput        = errors.New("error opening input file")
	ErrUnknownCharset   = errors.New("unknown charset")
	ErrParseInput       = errors.New("error parsing input file")
	ErrMarshalJSON      = errors.New("error marshalling json")
	ErrCreateOutputDir  = errors.New("error creating output directory")
	ErrCreateOutputFile = errors.New("error creating output file")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valute  []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type OutputItem struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := parseXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	sortValute(valCurs.Valute)

	outputItems, err := convertToOutput(valCurs.Valute)
	if err != nil {
		panic(err)
	}

	if err := saveOutput(config.OutputFile, outputItems); err != nil {
		panic(err)
	}
}

func loadConfig(path string) (Config, error) {
	var config Config

	configData, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("%w: %w", ErrReadConfig, err)
	}

	if err := yaml.Unmarshal(configData, &config); err != nil {
		return config, fmt.Errorf("%w: %w", ErrParseConfig, err)
	}

	return config, nil
}

func parseXML(path string) (ValCurs, error) {
	var valCurs ValCurs

	xmlFile, err := os.Open(path)
	if err != nil {
		return valCurs, fmt.Errorf("%w: %w", ErrOpenInput, err)
	}

	defer func() {
		_ = xmlFile.Close()
	}()

	decoder := xml.NewDecoder(xmlFile)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf("%w: %s", ErrUnknownCharset, charset)
	}

	if err := decoder.Decode(&valCurs); err != nil {
		return valCurs, fmt.Errorf("%w: %w", ErrParseInput, err)
	}

	return valCurs, nil
}

func parseValue(value string) (float64, error) {
	goodStringFloat := strings.ReplaceAll(value, ",", ".")

	res, err := strconv.ParseFloat(goodStringFloat, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrParseInput, err)
	}

	return res, nil
}

func sortValute(valutes []Valute) {
	sort.Slice(valutes, func(i, jIndex int) bool {
		goodFloat1, err := parseValue(valutes[i].Value)
		if err != nil {
			panic(err)
		}

		goodFloat2, err := parseValue(valutes[jIndex].Value)
		if err != nil {
			panic(err)
		}

		return goodFloat1 > goodFloat2
	})
}

func convertToOutput(valutes []Valute) ([]OutputItem, error) {
	outputItems := make([]OutputItem, 0, len(valutes))

	for _, val := range valutes {
		goodFloat, err := parseValue(val.Value)
		if err != nil {
			return nil, err
		}

		outputItems = append(outputItems, OutputItem{
			NumCode:  val.NumCode,
			CharCode: val.CharCode,
			Value:    goodFloat,
		})
	}

	return outputItems, nil
}

func saveOutput(path string, items []OutputItem) error {
	jsonData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrMarshalJSON, err)
	}

	outputDir := filepath.Dir(path)

	if err := os.MkdirAll(outputDir, outputDirPerm); err != nil {
		return fmt.Errorf("%w: %w", ErrCreateOutputDir, err)
	}

	if err := os.WriteFile(path, jsonData, outputFilePerm); err != nil {
		return fmt.Errorf("%w: %w", ErrCreateOutputFile, err)
	}

	return nil
}
