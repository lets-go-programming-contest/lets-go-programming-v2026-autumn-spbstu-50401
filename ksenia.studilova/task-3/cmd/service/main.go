package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
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
	outputDirMode  = 0o755
	outputFileMode = 0o600
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

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func convertCurrencies(valutes []Valute) []Currency {
	currencies := make([]Currency, 0, len(valutes))

	for _, valute := range valutes {
		valueStr := strings.ReplaceAll(valute.Value, ",", ".")

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			panic(err)
		}

		currency := Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		}

		currencies = append(currencies, currency)
	}

	return currencies
}

func higherValue(a, b Currency) bool {
	return a.Value > b.Value
}

func main() {
	configPath := flag.String("config", "", "path to YAML config file")
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		panic("input and output files must be specified")
	}

	xmlData, err := os.ReadFile(config.InputFile)
	if err != nil {
		panic(err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}

	currencies := convertCurrencies(valCurs.Valutes)

	sort.Slice(currencies, func(i, j int) bool {
		return higherValue(currencies[i], currencies[j])
	})

	jsonData, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(err)
	}

	outputDir := filepath.Dir(config.OutputFile)

	err = os.MkdirAll(outputDir, outputDirMode)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(config.OutputFile, jsonData, outputFileMode)
	if err != nil {
		panic(err)
	}
}
