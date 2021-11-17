package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fileReader(path string) []int {
	var numberArr []int
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numberArr = append(numberArr, value)
	}
	return numberArr
}

func partone(path string, preamble int) int {
	arr := fileReader(path)
	table := make(map[int]bool)
	for i := preamble + 1; i <= len(arr); i++ {
		nrrange := arr[i-(preamble+1) : i]
		nr := nrrange[len(nrrange)-1]
		var sumArr []int
		for j := 0; j < len(nrrange)-1; j++ {
			for k := 0; k < len(nrrange)-1; k++ {
				sumArr = append(sumArr, nrrange[j]+nrrange[k])
			}
		}
		for _, v := range sumArr {
			if v == nr {
				table[v] = true
			}
		}
	}
	for _, v := range arr[preamble:] {
		_, ok := table[v]
		if !ok {
			fmt.Println("Part one: ", v)
			return v
		}
	}
	return 0
}

func parttwo(path string, preamble int) {
	weakness := partone(path, preamble)
	arr := fileReader(path)
	var newarr []int
	for i, v := range arr {
		if v == weakness {
			newarr = arr[:i]
		}
	}
	var sum int
	var notfound bool
	var weaknessrange []int
	i := 0
	notfound = true
	for notfound {
		sum = 0
		for j := i; j < len(newarr)-1; j++ {
			sum += newarr[j]
			if sum == weakness {
				notfound = false
				weaknessrange = append(weaknessrange, i)
				weaknessrange = append(weaknessrange, j)
			}
		}
		i = i + 1
	}
	newrange := arr[weaknessrange[0] : weaknessrange[1]+1]
	var smallest, largest int
	smallest = 99999999
	for _, v := range newrange {
		if v < smallest {
			smallest = v
		} else if v > largest {
			largest = v
		}
	}
	fmt.Println("Part two: ", largest+smallest)
}

func main() {
	partone("input.txt", 25)
	parttwo("input.txt", 25)
}
