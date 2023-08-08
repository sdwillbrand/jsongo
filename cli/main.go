package main

import (
	"flag"
	"log"
	"os"

	"github.com/countersoda/jsongo"
)

func main() {
	filePath := flag.Arg(0)
	println(filePath)
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	res := jsongo.ParseFromByte(content)
	if res == nil {
		log.Fatal()
	}

}
