package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
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

	nm := regexp.MustCompile(`\s(\d){1,}`)
	im := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)`)

	instructions := [][]int{}
	stackSize := 0
	stackDepth := 0
	lineNumber := 0
	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}
		lineNumber++
		numLine := nm.FindAllStringSubmatch(line, -1)
		if stackSize == 0 && len(numLine) > 0 {
			last := numLine[len(numLine)-1]
			fmt.Sscan(last[1], &stackSize)
			stackDepth = lineNumber
		}
		instMatches := im.FindAllStringSubmatch(line, -1)
		if len(instMatches) > 0 {
			many, _ := strconv.ParseInt(instMatches[0][1], 0, 32)
			from, _ := strconv.ParseInt(instMatches[0][2], 0, 32)
			to, _ := strconv.ParseInt(instMatches[0][3], 0, 32)
			instructions = append(instructions, []int{int(many), int(from), int(to)})
		}
	}

	stacks := make([][]string, stackDepth)
	for _, line := range strings.Split(string(content), "\n") {
		// we've hit the numbers
		if len(nm.FindAllStringSubmatch(line, -1)) > 0 {
			break
		}
		// jump to where the chars are between the braces
		jumpCounter := 1
		for i := 0; i < stackSize; i++ {
			char := fmt.Sprintf("%c", line[jumpCounter])
			if char != " " {
				stacks[i] = append(stacks[i], char)
			}
			jumpCounter += 4
		}
	}

	for _, instruction := range instructions {
		instruction := instruction
		moves := instruction[0]
		from := instruction[1] - 1
		to := instruction[2] - 1

		toMove := []string{}
		for m := 0; m < moves; m++ {
			fromStack := stacks[from]
			x, fs := fromStack[0], fromStack[1:]
			toMove = append(toMove, x)
			stacks[from] = fs
		}

		for m := len(toMove) - 1; m >= 0; m-- {
			c := toMove[m]
			toStack := stacks[to]
			ts := append([]string{c}, toStack...)
			stacks[to] = ts
		}
	}
	fmt.Println(stacks)

	leading := []string{}
	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}
		leading = append(leading, stack[0])
	}
	fmt.Println(strings.Join(leading, ""))
}
