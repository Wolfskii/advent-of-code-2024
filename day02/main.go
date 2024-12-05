package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ReportLevelMovement string

const (
    NotStarted ReportLevelMovement = "NOT_STARTED"
    Increasing ReportLevelMovement = "INCREASING"
    Decreasing ReportLevelMovement = "DECREASING"
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
    defer file.Close()

    scanner := bufio.NewScanner(file)

    safeReports := 0

    // 1. Read the input file and get all reports (each line)
    for scanner.Scan() {
        reportLine := scanner.Text()

        // 2. Divide the reports's level numbers (variable amount) into an array of type [int]
        reportLevelsRaw := strings.Fields(reportLine)
        reportLevels := make([]int, len(reportLevelsRaw))

        for i, line := range reportLevelsRaw {
            num, err := strconv.Atoi(line)

            if err != nil {
                return 0, fmt.Errorf("invalid level found: %s", line)
            }

            reportLevels[i] = num
        }

        // 3. Validate that the levels (of each report) are either all increasing or all decreasing
        isLevelsMovementConsistent := validateConsistentLevelMovement(reportLevels)

        // If the levels are not consistent, skip the the next validation and skip the report since it's not safe
        if !isLevelsMovementConsistent {
            continue
        }

        // 4. Validate that any two adjacent levels (of each report) differ by at least one and at most three
        isAdjacentLevelsAcceptedRange := validateAdjacentLevelDifference(reportLevels, 1, 3)

        // 5. If both validations pass, increment the safeReports counter
        if (isLevelsMovementConsistent && isAdjacentLevelsAcceptedRange) {
            safeReports++
        }

        if err != nil {
            return 0, err
        }
    }

    if err := scanner.Err(); err != nil {
        return 0, err
    }

    return safeReports, nil
}

func validateConsistentLevelMovement(reportLevels []int) bool {
    if len(reportLevels) < 2 {
        return true
    }

    var reportLevelMovement ReportLevelMovement = NotStarted

    for i := 1; i < len(reportLevels); i++ {
        if reportLevels[i] < reportLevels[i-1] {
            if reportLevelMovement == NotStarted {
                reportLevelMovement = Decreasing
            } else if reportLevelMovement == Increasing {
                return false
            }
        } else if reportLevels[i] > reportLevels[i-1] {
            if reportLevelMovement == NotStarted {
                reportLevelMovement = Increasing
            } else if reportLevelMovement == Decreasing {
                return false
            }
        } else {
            // If two adjacent levels are equal, it's neither increasing nor decreasing
            return false
        }
    }

    return true
}

func validateAdjacentLevelDifference(reportLevels []int, smallestAllowedDiff int, biggestAllowedDiff int) (isAdjacentLevelsAcceptedRange bool) {
    for i := 1; i < len(reportLevels); i++ {
        adjacentDiff := reportLevels[i] - reportLevels[i-1]

        // Invert the difference if it's negative into a positive number
        if adjacentDiff < 0 {
            adjacentDiff = -adjacentDiff
        }

        if (adjacentDiff < smallestAllowedDiff || adjacentDiff > biggestAllowedDiff ) {
            fmt.Println("Report levels are not consistent")
            return false
        }
    }

    return true
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