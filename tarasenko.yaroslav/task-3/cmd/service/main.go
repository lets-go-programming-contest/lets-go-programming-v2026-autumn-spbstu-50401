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
	dirPerm  = 0o755
	filePerm = 0o600
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}
type OutputValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func loadConfig(filePath string) Config {
	flag.Parse()

	if filePath == "" {
		panic("config flag is required")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func parseXML(filePath string) ValCurs {
	inputData, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(inputData))
	decoder.CharsetReader = charset.NewReaderLabel

	var parsedData ValCurs

	err = decoder.Decode(&parsedData)
	if err != nil {
		panic(err)
	}

	return parsedData
}

func saveToJSON(filePath string, data []OutputValute) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(filePath)
	if dir != "" && dir != "." {
		err = os.MkdirAll(dir, dirPerm)
		if err != nil {
			panic(err)
		}
	}

	err = os.WriteFile(filePath, jsonData, filePerm)
	if err != nil {
		panic(err)
	}
}

func processValutes(valutes []Valute) []OutputValute {
	outputData := make([]OutputValute, 0, len(valutes))

	for _, val := range valutes {
		pointValue := strings.ReplaceAll(val.Value, ",", ".")

		floatValue, err := strconv.ParseFloat(pointValue, 64)
		if err != nil {
			panic(err)
		}

		outputData = append(outputData, OutputValute{
			NumCode:  val.NumCode,
			CharCode: val.CharCode,
			Value:    floatValue,
		})
	}

	sort.Slice(outputData, func(i, j int) bool {
		return outputData[i].Value > outputData[j].Value
	})

	return outputData
}

func main() {
	configPath := flag.String("config", "", "Path to YAML config")
	flag.Parse()

	if *configPath == "" {
		panic("config flag is required")
	}

	config := loadConfig(*configPath)
	parsedData := parseXML(config.InputFile)
	outputData := processValutes(parsedData.Valutes)
	saveToJSON(config.OutputFile, outputData)
}
