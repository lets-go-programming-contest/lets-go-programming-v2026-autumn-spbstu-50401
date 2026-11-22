package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
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

func main() {
	configPath := flag.String("config", "config.yaml", "путь к файлу конфигурации")
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

	xmlData, err := os.ReadFile(config.Input)
	if err != nil {
		panic(err)
	}

	var valCurs ValCurs

	err = xml.Unmarshal(xmlData, &valCurs)
	if err != nil {
		panic(err)
	}

	var currencies = make([]ValuteJSON, len(valCurs.Valutes))

	for _, v := range valCurs.Valutes {
		currency := ValuteJSON{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    parseValue(v.Value),
		}
		currencies = append(currencies, currency)
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	dataOut, err := json.MarshalIndent(currencies, "", " ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(config.Output, dataOut, 0600)
	if err != nil {
		panic(err)
	}
}
