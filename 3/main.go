package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var scores = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := string(input)
	rucksacks := strings.Split(content, "\n")

	score := 0
	for _, r := range rucksacks {
		seen := make(map[string]int)
		c1, c2 := arrangeRucksackCompartments(r)
		for _, c := range c1 {
			char := fmt.Sprintf("%c", c)
			count := strings.Count(c2, char)
			if count > 0 && seen[char] == 0 {
				score += scores[char]
				seen[char] = 1
			}
		}
	}
	fmt.Println("scores:", score)

	groups := formGroups(rucksacks)

	totalGroupScore := 0
	for _, g := range groups {
		// find common item in each rucksack
		rs1 := g[0]
		rs2 := g[1]
		rs3 := g[2]
		fmt.Println(rs1)
		fmt.Println(rs2)
		fmt.Println(rs3)
		groupScore := 0
		seen := make(map[string]int)
		for _, c := range rs1 {
			char := fmt.Sprintf("%c", c)
			found1 := strings.Count(rs2, char)
			found2 := strings.Count(rs3, char)
			if found1 > 0 && found2 > 0 && seen[char] == 0 {
				groupScore += scores[char]
				seen[char] = 1
			}
		}
		totalGroupScore += groupScore
		fmt.Println(groupScore)
	}
	fmt.Print(totalGroupScore)
}

func formGroups(r []string) [][]string {
	var groups [][]string
	size := len(r) / 3
	groupCounter := 0
	for i := 0; i < size; i++ {
		items := r[groupCounter : groupCounter+3]
		groupCounter += 3
		groups = append(groups, items)
	}
	return groups
}

func arrangeRucksackCompartments(r string) (c1 string, c2 string) {
	l := len(r)
	h := l / 2
	c1 = r[0:h]
	c2 = r[h:l]
	return
}
