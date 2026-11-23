package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/donkeev.grigoriy/task-3/config"
	"github.com/donkeev.grigoriy/task-3/parser"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o600
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	if flag.NArg() > 0 {
		panic("Too many arguments")
	}

	configVal, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Download error: %v", err))
	}

	xmlData, err := os.ReadFile(configVal.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Reading error: %v", err))
	}

	currencyData, err := parser.XMLFormatParser(xmlData)
	if err != nil {
		panic(fmt.Sprintf("Parsing to XML error: %v", err))
	}

	outputDir := filepath.Dir(configVal.OutputFile)
	if err := os.MkdirAll(outputDir, dirPermissions); err != nil {
		panic("Creating directory error")
	}

	jsonData, err := json.MarshalIndent(currencyData, "", "    ")
	if err != nil {
		panic("Coding to json error")
	}

	err = os.WriteFile(configVal.OutputFile, jsonData, filePermissions)
	if err != nil {
		panic("Error file writing")
	}
}
