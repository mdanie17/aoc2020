package main

import (
	"adventofcode/filereader"
	"fmt"
)

func firstpart() {
	numArray := filereader.FilereaderInt("input.txt")
	for _, v := range numArray {
		for _, v2 := range numArray {
			if v+v2 == 2020 {
				fmt.Printf("First part: %d + %d = %d | %d * %d = %d \n", v, v2, v+v2, v, v2, v*v2)
				return
			}
		}
	}
}

func secondpart() {
	numArray := filereader.FilereaderInt("input.txt")
	for _, v := range numArray {
		for _, v2 := range numArray {
			for _, v3 := range numArray {
				if v+v2+v3 == 2020 {
					fmt.Printf("Second part: %d + %d + %d = %d | %d * %d * %d = %d \n", v, v2, v3, v+v2+v3, v, v2, v3, v*v2*v3)
					return
				}
			}
		}
	}
}

func main() {
	firstpart()
	secondpart()
}
