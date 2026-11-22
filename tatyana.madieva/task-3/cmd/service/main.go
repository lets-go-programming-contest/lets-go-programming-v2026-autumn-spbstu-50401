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
	"golang.org/x/text/transform"
	"gopkg.in/yaml.v3"
)

const (
	dirPerm  = 0o755
	filePerm = 0o600
)

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

type ValCurs struct {
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func parseValue(s string) float64 {
	s = strings.Replace(s, ",", ".", 1)

	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return value
}

func readConfig(path string) Config {
	data, err := os.ReadFile(path)
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

func parseXML(path string) ValCurs {
	xmlData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var valCurs ValCurs

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return transform.NewReader(input, charmap.Windows1251.NewDecoder()), nil
		}

		return input, nil
	}

	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}

	return valCurs
}

func convertToJSON(valutes []Valute) []ValuteJSON {
	currencies := make([]ValuteJSON, 0, len(valutes))
	for _, v := range valutes {
		currencies = append(currencies, ValuteJSON{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    parseValue(v.Value),
		})
	}

	return currencies
}

func saveJSON(path string, data []ValuteJSON) {
	dataOut, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}

	outDir := filepath.Dir(path)

	err = os.MkdirAll(outDir, dirPerm)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path, dataOut, filePerm)
	if err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "config.yaml", "путь к файлу конфигурации")
	flag.Parse()

	config := readConfig(*configPath)
	valCurs := parseXML(config.Input)
	currencies := convertToJSON(valCurs.Valutes)

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	saveJSON(config.Output, currencies)
}
