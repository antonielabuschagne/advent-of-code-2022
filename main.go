package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	file := flag.String("file", "", "input filename")
	flag.Parse()
	if *file == "" {
		panic("file required argument")
	}
	content, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
