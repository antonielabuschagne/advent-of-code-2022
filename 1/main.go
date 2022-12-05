package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type ElfScore struct {
	score int
	num   int
}

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

	elfs := make(map[int]int)

	var elfNum int
	// batch the scores by elf
	for _, batch := range strings.Split(string(content), "\n\n") {
		var score int
		// count each of the individual scores
		for _, s := range strings.Split(batch, "\n") {
			if s == "" {
				continue
			}
			sc, err := strconv.ParseInt(s, 0, 64)
			if err != nil {
				panic(err)
			}
			score += int(sc)
		}
		elfNum++
		elfs[elfNum] = score
	}

	keys := make([]int, 0, len(elfs))
	for key := range elfs {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return elfs[keys[i]] < elfs[keys[j]]
	})

	first := keys[len(keys)-1]
	second := keys[len(keys)-2]
	third := keys[len(keys)-3]

	fmt.Printf("first elf %v with score %v\n", first, elfs[first])
	fmt.Printf("second elf %v with score %v\n", second, elfs[second])
	fmt.Printf("third elf %v with score %v\n", third, elfs[third])

	total := elfs[first] + elfs[second] + elfs[third]
	fmt.Printf("\n\n=====\nTOTAL: %v\n=====\n", total)
}
