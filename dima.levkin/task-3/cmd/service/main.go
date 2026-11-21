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

type Valute struct {
	XMLName   xml.Name `xml:"Valute"`
	ID        string   `xml:"ID,attr"`
	NumCode   int      `xml:"NumCode"`
	CharCode  string   `xml:"CharCode"`
	Nominal   int      `xml:"Nominal"`
	Name      string   `xml:"Name"`
	Value     string   `xml:"Value"`
	VunitRate string   `xml:"VunitRate"`
}

type Courses struct {
	XMLName    xml.Name `xml:"ValCurs"`
	ValuteList []Valute `xml:"Valute"`
}

type ResultCurrency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type Config struct {
	In  string `yaml:"input-file"`
	Out string `yaml:"output-file"`
}

func parseXML(data []byte) []ResultCurrency {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var courses Courses
	if err := decoder.Decode(&courses); err != nil {
		panic(err)
	}

	result := make([]ResultCurrency, 0, len(courses.ValuteList))

	for _, valute := range courses.ValuteList {
		valStr := strings.ReplaceAll(valute.Value, ",", ".")

		value, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			panic(err)
		}

		result = append(result, ResultCurrency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result
}

func writeOutput(path string, result []ResultCurrency) {
	resultBytes, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, resultBytes, filePerm); err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "", "path to config")
	flag.Parse()

	yml, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yml, &cfg); err != nil {
		panic(err)
	}

	inputFile, err := os.ReadFile(cfg.In)
	if err != nil {
		panic(err)
	}

	result := parseXML(inputFile)
	writeOutput(cfg.Out, result)
}
