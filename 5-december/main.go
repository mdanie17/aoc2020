package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rows    = 128
	columns = 8
)

func arrayBuilder(length int) []int {
	var array []int
	for i := 0; i < length; i++ {
		array = append(array, i)
	}
	return array
}

func fileReader(path string) []string {
	var boardingpassList []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boardingpassList = append(boardingpassList, scanner.Text())
	}
	return boardingpassList
}

func rowFinder(boardingpass string) int {
	var guess []int
	arr := arrayBuilder(rows)
	rowstring := boardingpass[:7]
	guess = arr
	for _, v := range rowstring {
		if v == 'F' {
			guess = guess[:len(guess)/2]
		}
		if v == 'B' {
			guess = guess[len(guess)/2:]
		}
	}
	return guess[0]
}

func columnFinder(boardingpass string) int {
	var guess []int
	arr := arrayBuilder(columns)
	columnstring := boardingpass[len(boardingpass)-3:]
	guess = arr
	for _, v := range columnstring {
		if v == 'L' {
			guess = guess[:len(guess)/2]
		}
		if v == 'R' {
			guess = guess[len(guess)/2:]
		}
	}
	return guess[0]
}

func partone(path string) {
	var highest int = 0
	list := fileReader(path)

	for _, v := range list {
		row := rowFinder(v)
		col := columnFinder(v)
		seatId := row*8 + col
		if seatId > highest {
			highest = seatId
		}
	}
	fmt.Println("First exercise: ", highest)
}

func parttwo(path string) {
	seatIdList := make(map[int]bool)
	list := fileReader(path)
	for _, v := range list {
		row := rowFinder(v)
		col := columnFinder(v)
		seatId := row*8 + col
		seatIdList[seatId] = true
	}
	for i := 0; i < len(seatIdList); i++ {
		_, ok := seatIdList[i]
		if !ok {
			_, ok = seatIdList[i-1]
			if ok {
				_, ok = seatIdList[i+1]
				if ok {
					fmt.Println("Second exercise: ", i)
				}
			}
		}
	}
}

func main() {
	partone("input.txt")
	parttwo("input.txt")
}
