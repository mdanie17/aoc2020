package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func fileReader(path string) []int {
	var values []int
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valueInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, valueInt)
	}
	sort.Ints(values)
	return values
}

func checkArray(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func partone(path string) {
	arr := fileReader(path)
	var currentValue int = 0
	var onecounter, threecounter int
	for _, v := range arr {
		if v-currentValue == 1 {
			currentValue = v
			onecounter++
		} else if v-currentValue == 2 {
			currentValue = v
		} else if v-currentValue == 3 {
			currentValue = v
			threecounter++
		}
	}
	fmt.Println("Answer: ", onecounter*(threecounter+1))
}

func count_paths_to_end(arr []int, table map[int]bool, val int) int {
	var sum int
	if val == arr[len(arr)-1] {
		return 1
	}
	_, ok := table[val]
	if ok {
		for i := 0; i < 3; i++ {
			sum += count_paths_to_end(arr, table, val+i+1)
		}
		return sum
	}
	return 0
}

func parttwo(path string) {
	arr := fileReader(path)
	arr = append([]int{0}, arr...)
	arr = append(arr, arr[len(arr)-1]+3)
	table := make(map[int]bool)
	for _, v := range arr {
		table[v] = true
	}
	answer := count_paths_to_end(arr, table, 0)
	fmt.Println("Part two: ", answer)
}

func main() {
	partone("input.txt")
	parttwo("input.txt")
}
