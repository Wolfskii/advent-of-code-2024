package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
    
    totalDistance := 0
    leftNumbers := []int{}
    rightNumbers := []int{}

    // 1. Read the input file and store the left and right numbers in separate slices
    for scanner.Scan() {
        line := scanner.Text()

        
        var left, right int
        _, err := fmt.Sscanf(line, "%d %d", &left, &right)
        if err != nil {
            return 0, err
        }

        leftNumbers = append(leftNumbers, left)
        rightNumbers = append(rightNumbers, right)
    }

    // 2. Sort left and right numbers by size from smallest first
    slices.Sort(leftNumbers)
    slices.Sort(rightNumbers)

    // 3. Calculate the distance between left and right numbers, based on size, starting with the pair of smallest of each side, then the pair of second smallest and so on
    for i := 0; i < len(leftNumbers); i++ {
        distance := leftNumbers[i] - rightNumbers[i]

        // Inverse negative differences to positive numbers
        if distance < 0 {
            distance = -distance
        }

        // 4. Add the distance to the total distance
        totalDistance += distance
    }

    if err := scanner.Err(); err != nil {
        return 0, err
    }

    return totalDistance, nil
}

func solvePart2(inputPath string) (int, error) {
    file, err := os.Open(inputPath)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    
    similarityScore := 0
    leftNumbers := []int{}
    rightNumbers := []int{}

    // 1. Read the input file and store the left and right numbers in separate slices
    for scanner.Scan() {
        line := scanner.Text()
        
        var left, right int
        _, err := fmt.Sscanf(line, "%d %d", &left, &right)
        if err != nil {
            return 0, err
        }

        leftNumbers = append(leftNumbers, left)
        rightNumbers = append(rightNumbers, right)
    }

    // 2. Check for each number in the leftNumbers how many times it appears in the rightNumbers
    for i := 0; i < len(leftNumbers); i++ {
        totalAppearances := 0

        for j := 0; j < len(rightNumbers); j++ {
            if (leftNumbers[i] == rightNumbers[j]) {
                totalAppearances++
            }
        }

        // 3. Multiply the number of appearances by the number itself and add it to the similarity score
        similarityScore += leftNumbers[i] * totalAppearances
    }

    if err := scanner.Err(); err != nil {
        return 0, err
    }

    return similarityScore, nil
}