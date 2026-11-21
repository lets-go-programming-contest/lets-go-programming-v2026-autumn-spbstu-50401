package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"golang.org/x/net/html/charset"
)

var (
	ErrReadingConfigFile = errors.New("error reading config file")
	ErrReadingXML        = errors.New("error reading XML file")
)

type Valute struct {
	XMLName   xml.Name `xml:"Valute"`
	Id        string   `xml:"ID,attr"`
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

func main() {
	path_ptr := flag.String("config", "", "path to config")
	flag.Parse()

	yml, err := os.ReadFile(*path_ptr)
	if err != nil {
		panic(fmt.Sprint(ErrReadingConfigFile))
	}

	var cfg struct {
		In  string `yaml:"input-file"`
		Out string `yaml:"output-file"`
	}
	if err := yaml.Unmarshal([]byte(yml), &cfg); err != nil {
		panic(err)
	}

	input_file, err := os.ReadFile(cfg.In)
	if err != nil {
		panic(fmt.Sprint(ErrReadingXML))
	}
	decoder := xml.NewDecoder(bytes.NewReader(input_file))
	decoder.CharsetReader = charset.NewReaderLabel
	var b Courses
	if err := decoder.Decode(&b);err != nil {
		panic(err)
	}
	fmt.Print(b)

}
