package main

import (
	"encoding/json"
	"encoding/xml"
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

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	Items []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "Path to YAML config file")
	flag.Parse()

	if *configPath == "" {
		panic("The -config flag is required")
	}

	config := readConfig(*configPath)
	fmt.Println("Config loaded successfully")

	currencies := parseXML(config.InputFile)
	fmt.Printf("Parsed %d currencies\n", len(currencies))

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	saveJSON(config.OutputFile, currencies)
	fmt.Println("Success! Result saved to:", config.OutputFile)
}

func readConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("Failed to read config file: " + err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic("Failed to parse YAML config: " + err.Error())
	}

	return cfg
}

func parseXML(path string) []CurrencyJSON {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("Failed to read XML file: " + err.Error())
	}

	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return input, nil
	}

	var valCurs ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		panic("Failed to parse XML: " + err.Error())
	}

	var result []CurrencyJSON
	for _, item := range valCurs.Items {
		code, _ := strconv.Atoi(strings.TrimSpace(item.NumCode))

		valStr := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			panic("Invalid Value format: " + item.Value)
		}

		result = append(result, CurrencyJSON{
			NumCode:  code,
			CharCode: item.CharCode,
			Value:    val,
		})
	}

	return result
}

func saveJSON(path string, data []CurrencyJSON) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic("Failed to serialize JSON: " + err.Error())
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic("Failed to create output directory: " + err.Error())
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		panic("Failed to write output file: " + err.Error())
	}
}
