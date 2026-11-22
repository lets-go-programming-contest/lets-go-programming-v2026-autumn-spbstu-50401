package main

import (
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
	dirPerm  = 0o755
	filePerm = 0o600
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
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type OutputCurrency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func LoadConfig(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}

func ParseXML(path string) []Valute {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	var valCurs ValCurs

	if err := decoder.Decode(&valCurs); err != nil {
		panic(err)
	}

	return valCurs.Valutes
}

func parseValue(s string) float64 {
	s = strings.ReplaceAll(s, ",", ".")

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return f
}

func sortByValueDesc(valutes []Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return parseValue(valutes[i].Value) > parseValue(valutes[j].Value)
	})
}

func SaveJSON(path string, valutes []Valute) {
	out := make([]OutputCurrency, 0, len(valutes))

	for _, valute := range valutes {
		num, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			panic(err)
		}

		out = append(out, OutputCurrency{
			NumCode:  num,
			CharCode: valute.CharCode,
			Value:    parseValue(valute.Value),
		})
	}

	data, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		panic(err)
	}

	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, dirPerm); err != nil {
			panic(err)
		}
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "", "path to yaml config")
	flag.Parse()

	if *configPath == "" {
		panic("config path is required (-config)")
	}

	cfg := LoadConfig(*configPath)
	valutes := ParseXML(cfg.InputFile)
	sortByValueDesc(valutes)
	SaveJSON(cfg.OutputFile, valutes)
}
