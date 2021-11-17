package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type coordinates struct {
	x int
	y int
}

func fileReader(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func adjacentChecker(arr []string, x, y int) int {
	var counter int

	//Middle of grid
	if (x != 0 && x != len(arr[y])-1) && (y != 0 && y != len(arr)-1) {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i+y == y && x+j == x {
					continue
				}
				if string(arr[y+i][x+j]) == "#" {
					counter++
				}
			}
		}
		//Middle of grid, left end
	} else if x == 0 && y != 0 && y != len(arr)-1 {
		for i := -1; i < 2; i++ {
			for j := 0; j < 2; j++ {
				if i+y == y && x+j == x {
					continue
				}
				if string(arr[y+i][x+j]) == "#" {
					counter++
				}
			}
		}
		//Middle of grid, right end
	} else if x == len(arr[y])-1 && y != 0 && y != len(arr)-1 {
		for i := -1; i < 2; i++ {
			for j := -1; j < 1; j++ {
				if i+y == y && x+j == x {
					continue
				}
				if string(arr[y+i][x+j]) == "#" {
					counter++
				}
			}
		}
		//Top line, middle
	} else if y == 0 && x != 0 && x != len(arr[y])-1 {
		for i := 0; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i+y == y && x+j == x {
					continue
				}
				if string(arr[y+i][x+j]) == "#" {
					counter++
				}
			}

		}
		//Top line, right end
	} else if y == 0 && x == len(arr[y])-1 {
		for i := 0; i < 2; i++ {
			for j := -1; j < 1; j++ {
				if i+y == y && x+j == x {
					continue
				}
				if string(arr[y+i][x+j]) == "#" {
					counter++
				}
			}
		}
		//Top line, left end
	} else if y == 0 && x == 0 {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				if i+y == y && x+j == x {
					continue
				}
				if string(arr[y+i][x+j]) == "#" {
					counter++
				}
			}
		}
		//Bottomline, left end
	} else if y == len(arr)-1 && x == 0 {
		for i := -1; i < 1; i++ {
			for j := 0; j < 2; j++ {
				if i+y == y && x+j == x {
					continue
				}
				if string(arr[y+i][x+j]) == "#" {
					counter++
				}
			}
		}
		//Bottomline, middle
	} else if y == len(arr)-1 && x != 0 && x != len(arr[y])-1 {
		for i := -1; i < 1; i++ {
			for j := -1; j < 2; j++ {
				if i+y == y && x+j == x {
					continue
				}
				if string(arr[y+i][x+j]) == "#" {
					counter++
				}
			}
		}
	} else if y == len(arr)-1 && x == len(arr[y])-1 {
		for i := -1; i < 1; i++ {
			for j := -1; j < 1; j++ {
				if i+y == y && x+j == x {
					continue
				}
				if string(arr[y+i][x+j]) == "#" {
					counter++
				}
			}
		}
	}
	return counter
}

func occupyCounter(arr []string) int {
	var counter int
	for _, v := range arr {
		for _, seat := range v {
			if string(seat) == "#" {
				counter++
			}
		}
	}
	return counter
}

func firstRound(arr []string) {
	for i, v := range arr {
		arr[i] = strings.Replace(v, "L", "#", -1)
	}
}

func ruleTwo(arr []string) {
	var seatList []coordinates
	for i, v := range arr {
		for j, k := range v {
			if k == '.' {
				continue
			}
			count := adjacentChecker(arr, j, i)
			if count >= 4 {
				seatList = append(seatList, coordinates{x: j, y: i})
				//arr[i] = arr[i][:j] + "L" + arr[i][j+1:]
			}
		}
	}
	for _, v := range seatList {
		arr[v.y] = arr[v.y][:v.x] + "L" + arr[v.y][v.x+1:]
	}
	//fmt.Println(seatList)
}

func ruleTwoSit(arr []string) {
	var seatList []coordinates

	for i, v := range arr {
		for j, k := range v {
			if k == '.' || k == '#' {
				continue
			}
			count := adjacentChecker(arr, j, i)
			if count == 0 {
				seatList = append(seatList, coordinates{x: j, y: i})
			}
		}
	}
	for _, v := range seatList {
		arr[v.y] = arr[v.y][:v.x] + "#" + arr[v.y][v.x+1:]
	}
	//fmt.Println(seatList)
}

func arrayPrinter(arr []string) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func partone(path string) {
	arr := fileReader(path)
	var lastcount int
	firstRound(arr)
	count := occupyCounter(arr)
	for count != lastcount {
		ruleTwo(arr)
		ruleTwoSit(arr)
		ruleTwo(arr)
		lastcount = count
		count = occupyCounter(arr)
	}
	fmt.Println("Part one: ", count)
}

func main() {
	partone("input.txt")
}
