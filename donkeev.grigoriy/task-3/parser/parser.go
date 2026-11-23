package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

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

func XMLFormatParser(data []byte) ([]OutputCurrency, error) {
	var valCurs ValCurs

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("xml decode: %w", err)
	}

	result := make([]OutputCurrency, 0, len(valCurs.Valutes))

	for _, value := range valCurs.Valutes {
		var numCode int

		if value.NumCode != "" {
			parsedCode, err := strconv.Atoi(value.NumCode)
			if err != nil {
				return nil, fmt.Errorf("types conversion: %w", err)
			}

			numCode = parsedCode
		}

		valueStr := strings.ReplaceAll(value.Value, ",", ".")

		valueFloat, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("types conversion: %w", err)
		}

		result = append(result, OutputCurrency{
			NumCode:  numCode,
			CharCode: value.CharCode,
			Value:    valueFloat,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result, nil
}
