package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    inputPath := "input.txt"

    part1, part2, err := Solve(inputPath)
    if err != nil {
        log.Fatalf("Error solving puzzle: %v", err)
    }

    fmt.Printf("Part 1 Result: %d\n", part1)
    fmt.Printf("Part 2 Result: %d\n", part2)
}

// Solve runs the solution for the day's puzzle
func Solve(inputPath string) (int, int, error) {
    // Part 1 solution
    part1Result, err := solvePart1(inputPath)
    if err != nil {
        return 0, 0, err
    }

    // Part 2 solution
    part2Result, err := solvePart2(inputPath)
    if err != nil {
        return 0, 0, err
    }

    return part1Result, part2Result, nil
}

func solvePart1(inputPath string) (int, error) {
    file, err := os.Open(inputPath)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // 1. Read the input file and get all reports (each line)
    for scanner.Scan() {
        reportLine := scanner.Text()

        
        // 2. Divide the reports's five level numbers directly into an array
        report := [5]int{}

        _, err := fmt.Sscanf(reportLine, "%d %d %d %d %d", &report[0], &report[1], &report[2], &report[3], &report[4])

        fmt.Println(report)    

        if err != nil {
            return 0, err
        }
    }

    if err := scanner.Err(); err != nil {
        return 0, err
    }

    return 0, nil
}

func solvePart2(inputPath string) (int, error) {
    file, err := os.Open(inputPath)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    sum := 0

/*     for scanner.Scan() {
        line := scanner.Text()

    } */

    if err := scanner.Err(); err != nil {
        return 0, err
    }

    return sum, nil
}