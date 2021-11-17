package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mapReader(path string) []string {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func lineRepeater(lines []string, slope int) []string {
	var linesRepeated []string
	for _, v := range lines {
		linesRepeated = append(linesRepeated, strings.Repeat(v, slope))
	}
	return linesRepeated
}

func treeChecker(lines []string, slope, descend int) int {
	var counter int
	var positiony = 0
	var positionx = 0

	for {
		positiony = positiony + descend
		positionx = positionx + slope

		if positiony > len(lines)-1 {
			break
		}

		if string(lines[positiony][positionx]) == "#" {
			counter = counter + 1
		}

	}
	return counter
}

func firstpart() {
	slope := 3
	descend := 1

	initialmap := mapReader("input.txt")
	multiplier := (len(initialmap) * slope) + 1
	puzzlemap := lineRepeater(initialmap, multiplier)
	answer := treeChecker(puzzlemap, slope, descend)
	fmt.Println("First exercise:	", answer)
}

func secondpart() {
	slopes := []int{1, 3, 5, 7, 1}
	descends := []int{1, 1, 1, 1, 2}

	initialmap := mapReader("input.txt")
	var answer int = 1
	for i, _ := range slopes {
		multiplier := (len(initialmap) * slopes[i]) + 1
		puzzlemap := lineRepeater(initialmap, multiplier)
		answer = answer * treeChecker(puzzlemap, slopes[i], descends[i])
	}
	fmt.Println("Second exercise:	", answer)
}

func main() {
	firstpart()
	secondpart()
}
