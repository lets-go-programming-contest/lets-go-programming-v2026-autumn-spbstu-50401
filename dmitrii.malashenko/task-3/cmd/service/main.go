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

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			panic(closeErr)
		}
	}()

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

func parseValue(s string) (float64, bool) {
	s = strings.ReplaceAll(strings.TrimSpace(s), ",", ".")

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, false
	}

	return f, true
}

func toOutputCurrency(valute Valute) (OutputCurrency, bool) {
	numCode := strings.TrimSpace(valute.NumCode)
	charCode := strings.TrimSpace(valute.CharCode)

	if numCode == "" || charCode == "" {
		return OutputCurrency{}, false
	}

	num, err := strconv.Atoi(numCode)
	if err != nil {
		return OutputCurrency{}, false
	}

	value, ok := parseValue(valute.Value)
	if !ok {
		return OutputCurrency{}, false
	}

	return OutputCurrency{
		NumCode:  num,
		CharCode: charCode,
		Value:    value,
	}, true
}

func sortByValueDesc(valutes []Valute) {
	sort.SliceStable(valutes, func(i, j int) bool {
		valueI, _ := parseValue(valutes[i].Value)
		valueJ, _ := parseValue(valutes[j].Value)

		return valueI > valueJ
	})
}

func SaveJSON(path string, valutes []Valute) {
	out := make([]OutputCurrency, 0, len(valutes))

	for _, valute := range valutes {
		item, ok := toOutputCurrency(valute)
		if !ok {
			continue
		}

		out = append(out, item)
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
