package main

import (
	"flag"
	"os"

	"github.com/countersoda/jsongo"
)

func main() {
	flag.Parse()
	filePath := flag.Arg(0)
	content, err := os.ReadFile(filePath)
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}()
	if err != nil {
		os.Exit(1)
	}
	res := jsongo.ParseFromByte(content)
	if res == nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
