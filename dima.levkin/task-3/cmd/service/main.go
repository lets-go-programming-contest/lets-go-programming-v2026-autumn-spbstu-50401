package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

var (
	ErrReadingConfigFile = errors.New("error reading config file")
	ErrParsingYaml       = errors.New("error parsing YAML")
)

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
	fmt.Println(cfg)

}
