package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	ErrReadingConfigFile = errors.New("error reading config file")
)

func main() {
	path_ptr := flag.String("config", "", "path to config")
	flag.Parse()

	data, err := os.ReadFile(*path_ptr)
	if err != nil {
		panic(fmt.Sprint(ErrReadingConfigFile))
	}



}
