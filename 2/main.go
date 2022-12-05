package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var opponentHandDef = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
}

var playerHandDef = map[string]string{
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

var opponentToPlayerHand = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var handScoreDef = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var handOutcomes = map[string]string{
	"X": "lose",
	"Y": "draw",
	"Z": "win",
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

	winningBonus := 6
	drawBonus := 3
	totalScore := 0
	for _, round := range strings.Split(string(content), "\n") {
		if round == "" {
			continue
		}
		hands := strings.Split(round, " ")
		oHandRaw := hands[0]
		// step 1, where the hand is based on the file
		// pHandRaw := hands[1]
		pHandRaw := choosePlayerHand(oHandRaw, handOutcomes[hands[1]])
		oHand := opponentHandDef[oHandRaw]
		pHand := playerHandDef[pHandRaw]
		//pHand := choosePlayerHand(hands[0], handOutcomes[hands[1]])
		pScore := handScoreDef[pHandRaw]
		// it's a draw
		if oHand == pHand {
			pScore += drawBonus
		} else if (pHand == "Rock" && oHand == "Scissors") || (pHand == "Scissors" && oHand == "Paper") || (pHand == "Paper" && oHand == "Rock") {
			pScore += winningBonus
		}
		totalScore += pScore
	}

	fmt.Printf("Player Score: %v\n", totalScore)
}

func choosePlayerHand(o string, outcome string) string {
	if outcome == "draw" {
		return opponentToPlayerHand[o]
	}
	oHand := opponentHandDef[o]
	if outcome == "lose" {
		switch oHand {
		case "Rock":
			// scissors
			return "Z"
		case "Scissors":
			// paper
			return "Y"
		case "Paper":
			// rock
			return "X"
		}
	}

	// we want to win :)
	switch oHand {
	case "Rock":
		// paper
		return "Y"
	case "Scissors":
		// rock
		return "X"
	case "Paper":
		// scissors
		return "Z"
	default:
		panic("unknown hand")
	}
}
