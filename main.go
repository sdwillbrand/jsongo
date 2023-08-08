package main

import (
	"flag"
	"fmt"
	jq "jsongo/pkg"
	"log"
	"os"
)

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	res := jq.ParseFromByte(data)
	if res == nil {
		log.Fatal("Invalid json")
	}
	fmt.Print(res.String())
}
