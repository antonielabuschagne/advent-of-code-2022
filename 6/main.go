package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	message := flag.String("message", "", "packet/message")
	size := flag.Int("marker-size", 0, "number or characters to identify marker")
	flag.Parse()
	if *message == "" || *size <= 0 {
		panic("message & marker-size required")
	}

	ds := *message
	ms := *size
	for i := 0; i < len(ds)-1; i++ {
		max := i + ms
		if max > len(ds) {
			break
		}
		batch := ds[i:max]
		if isStartOfPacket(batch, ms) {
			fmt.Println(strings.Index(ds, batch) + ms)
			break
		}
	}
}

func isStartOfPacket(s string, ms int) bool {
	seen := make(map[string]int)
	for _, u := range s {
		char := fmt.Sprintf("%c", u)
		seen[char]++
	}
	return len(seen) == ms
}
