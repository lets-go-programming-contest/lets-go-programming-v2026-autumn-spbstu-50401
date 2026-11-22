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

	"go.yaml.in/yaml/v4"
	"golang.org/x/text/encoding/charmap"
)

type Float float64

func (f *Float) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string

	err := d.DecodeElement(&str, &start)
	if err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	str = strings.ReplaceAll(strings.TrimSpace(str), ",", ".")
	if str == "" {
		return nil
	}

	f64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("parse float: %w", err)
	}

	*f = Float(f64)

	return nil
}

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName     xml.Name `xml:"ValCurs"`
	Date        string   `xml:"Date,attr"`
	Name        string   `xml:"name,attr"`
	ValuteItems []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName   xml.Name `xml:"Valute"`
	ID        string   `xml:"ID,attr"`
	NumCode   int      `xml:"NumCode"`
	CharCode  string   `xml:"CharCode"`
	Nominal   int      `xml:"Nominal"`
	Name      string   `xml:"Name"`
	Value     Float    `xml:"Value"`
	VunitRate Float    `xml:"VunitRate"`
}

type ValuteOutput struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

var errUnsupportedCharset = errors.New("unsupported charset")

const dirPerm = 0o755

func newDecoder(input io.Reader) *xml.Decoder {
	decoder := xml.NewDecoder(input)
	decoder.CharsetReader = func(charset string, r io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(r), nil
		}

		return nil, fmt.Errorf("%w: %q", errUnsupportedCharset, charset)
	}

	return decoder
}

func readXML(path string) ValCurs {
	xmlFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			panic(err)
		}
	}(xmlFile)

	var valCurs ValCurs

	decoder := newDecoder(xmlFile)
	if err := decoder.Decode(&valCurs); err != nil {
		panic(err)
	}

	return valCurs
}

func getConfig() Config {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	configData, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func writeJSON(path string, data []ValuteOutput) {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, dirPerm)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(data)
	if err != nil {
		panic(err)
	}
}

func createValuteOutput(valCurs *ValCurs) []ValuteOutput {
	result := make([]ValuteOutput, 0, len(valCurs.ValuteItems))

	for _, v := range valCurs.ValuteItems {
		item := ValuteOutput{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    float64(v.Value),
		}
		result = append(result, item)
	}

	return result
}

func main() {
	config := getConfig()

	valCurs := readXML(config.InputFile)

	sort.Slice(valCurs.ValuteItems, func(i, j int) bool {
		return valCurs.ValuteItems[i].Value > valCurs.ValuteItems[j].Value
	})

	valuteOutput := createValuteOutput(&valCurs)
	writeJSON(config.OutputFile, valuteOutput)
}
