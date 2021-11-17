package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type color struct {
	Amt   uint
	Color string
}

func main() {
	rules := parseRules("input.txt")
	fmt.Printf("Part 1: %d\n", part1(rules))
	fmt.Printf("Part 2: %d\n", part2(rules))
}

// parseRules returns a map of mainColors mapped to the rest of the rules colors
func parseRules(path string) map[string][]color {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	rules := make(map[string][]color)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split rule on space
		parts := strings.Split(scanner.Text(), " ")

		// Get main color
		mainColor := strings.Join(parts[:2], " ")

		// Prepare rest of colors slice
		rest := make([]color, 0)

		i := 5
		for {
			// At the end of the rule
			if len(parts) < i+2 {
				break
			}

			// Account for "no other" which will just be 0
			amt, err := strconv.Atoi(parts[i-1])
			if err != nil {
				amt = 0
			}

			rest = append(rest, color{
				Color: strings.Join(parts[i:i+2], " "),
				Amt:   uint(amt),
			})

			// Next color is 4 indexes ahead if there is one
			i += 4
		}

		rules[mainColor] = rest
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rules
}

func part1(rules map[string][]color) uint {
	// Set of bags that can contain a shiny gold bag
	contains := make(map[string]struct{})
	check("shiny gold", rules, contains)
	return uint(len(contains))
}

func check(color string, rules map[string][]color, contains map[string]struct{}) {
	for c := range getContains(rules, color) {
		contains[c] = struct{}{}
		check(c, rules, contains)
	}
}

// Returns all the bags that contain the given bag
func getContains(input map[string][]color, bag string) map[string]struct{} {
	contains := make(map[string]struct{})
	for mainColor, restColors := range input {
		for _, clr := range restColors {
			if clr.Color == bag {
				contains[mainColor] = struct{}{}
			}
		}
	}
	return contains
}

func part2(rules map[string][]color) uint {
	return getAmt("shiny gold", rules)
}

func getAmt(color string, rules map[string][]color) uint {
	var amt uint
	for _, clr := range rules[color] {
		amt += clr.Amt
		amt += clr.Amt * getAmt(clr.Color, rules)
	}
	return amt
}
