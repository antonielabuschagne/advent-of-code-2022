package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	c, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	sections := string(c)

	full := 0
	partial := 0
	for _, s := range strings.Split(sections, "\n") {
		elfSections := strings.Split(s, ",")

		if len(elfSections) < 2 {
			continue
		}
		min1, max1 := getBoundaries(elfSections[0])
		min2, max2 := getBoundaries(elfSections[1])

		sec1Map := expandRange(min1, max1)
		sec2Map := expandRange(min2, max2)

		if (min1 < min2 && max1 > max2) || (min2 < min1 && max2 > max1) {
			full += 1
			continue
		}

		found := false
		for _, n := range sec1Map {
			if _, ok := sec2Map[n]; ok {
				found = true
			}
		}
		if found {
			partial++
		}
	}
	fmt.Printf("full: %v\n", full)
	fmt.Printf("partial: %v\n", partial)
	fmt.Printf("total: %v\n", full+partial)
}

func expandRange(min int64, max int64) map[int64]int64 {
	ran := make(map[int64]int64)
	for i := min; i <= max; i++ {
		ran[i] = i
	}
	return ran
}

func getBoundaries(r string) (min int64, max int64) {
	secRange := strings.Split(r, "-")
	if len(secRange) < 2 {
		panic(fmt.Sprintf("invalid format for range: %s", r))
	}
	min, _ = strconv.ParseInt(secRange[0], 0, 32)
	max, _ = strconv.ParseInt(secRange[1], 0, 32)
	return
}
