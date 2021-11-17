package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func firstpart() {
	filepath := "input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	csv := csv.NewReader(file)

	csv.Comma = ' '
	var bounds, letter, sequence string
	var counter int
	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		bounds = record[0]
		boundssplit := strings.Split(bounds, "-")
		letter = strings.Trim(record[1], ":")
		sequence = record[2]

		upperbound, _ := strconv.Atoi(string(boundssplit[1]))
		lowerbound, _ := strconv.Atoi(string(boundssplit[0]))

		if (lowerbound <= strings.Count(sequence, letter)) && (strings.Count(sequence, letter) <= upperbound) {
			counter = counter + 1
		}
	}
	fmt.Println(counter)

}

func secondpart() {
	filepath := "input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	csv := csv.NewReader(file)

	csv.Comma = ' '
	var bounds, letter, sequence string
	var counter int
	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		bounds = record[0]
		boundssplit := strings.Split(bounds, "-")
		letter = strings.Trim(record[1], ":")
		sequence = record[2]

		index1, _ := strconv.Atoi(string(boundssplit[1]))
		index2, _ := strconv.Atoi(string(boundssplit[0]))

		if string(sequence[index1-1]) == letter && string(sequence[index2-1]) != letter || string(sequence[index2-1]) == letter && string(sequence[index1-1]) != letter {
			counter = counter + 1
		}
	}
	fmt.Println(counter)

}

func main() {
	firstpart()
	secondpart()
}
