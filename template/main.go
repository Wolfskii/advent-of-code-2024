package main

import (
	"fmt"
	"log"

	"advent-of-code-2024/dayXX"
)

func main() {
	inputPath := "input.txt"

	part1, part2, err := dayXX.Solve(inputPath)
	if err != nil {
		log.Fatalf("Error solving puzzle: %v", err)
	}

	fmt.Printf("Part 1 Result: %d\n", part1)
	fmt.Printf("Part 2 Result: %d\n", part2)
}