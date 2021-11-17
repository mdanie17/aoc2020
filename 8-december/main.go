package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instructionStruct struct {
	instruction string
	value       int
	counter     int
}

func fileReader(path string) []string {
	var instructions []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	return instructions
}

func instructionParser(instructions []string) []instructionStruct {
	var instructionsStructs []instructionStruct
	for _, v := range instructions {
		splitted := strings.Split(v, " ")
		valueint, err := strconv.Atoi(splitted[1])
		if err != nil {
			fmt.Println(err, splitted[1])
		}
		ins := instructionStruct{
			instruction: splitted[0],
			value:       valueint,
		}
		instructionsStructs = append(instructionsStructs, ins)
	}
	return instructionsStructs
}

func instructionHandler(instructions []instructionStruct) (int, bool) {
	var accumulator int
	for i := 0; i < len(instructions); i++ {
		switch instructions[i].instruction {
		case "nop":
			//fmt.Println("Nop", instructions[i].instruction, instructions[i].counter)
		case "acc":
			//fmt.Println("Acc", instructions[i].value, instructions[i].counter)
			accumulator += instructions[i].value
			//fmt.Println(instructions[i].value)
		case "jmp":
			//fmt.Println("Jmp", instructions[i].instruction, instructions[i].counter)
			i += instructions[i].value - 1
		}
		if instructions[i].counter > 0 {
			return accumulator, false
		} else {
			instructions[i].counter += 1
		}
		if i == len(instructions)-1 {
			return accumulator, true
		}
	}

	return accumulator, false
}

func jmpFinder(arr []instructionStruct) []int {
	var indexArr []int
	for i, v := range arr {
		if v.instruction == "jmp" {
			indexArr = append(indexArr, i)
		}
	}
	return indexArr
}

func nopFinder(arr []instructionStruct) []int {
	var indexArr []int
	for i, v := range arr {
		if v.instruction == "nop" {
			indexArr = append(indexArr, i)
		}
	}
	return indexArr
}

func partone(path string) {
	sample := fileReader(path)
	arr := instructionParser(sample)
	value, end := instructionHandler(arr)
	fmt.Printf("Partone: Reached end:	%v   with value:	%d\n", end, value)
}

func counterReset(arr []instructionStruct) []instructionStruct {
	for i, _ := range arr {
		arr[i].counter = 0
	}
	return arr
}

func parttwo(path string) {
	sample := fileReader(path)
	arr := instructionParser(sample)
	jmpIndex := jmpFinder(arr)
	nopIndex := nopFinder(arr)

	for i := 0; i < len(jmpIndex); i++ {
		arr := counterReset(arr)
		arr[jmpIndex[i]].instruction = "nop"
		value, end := instructionHandler(arr)
		if !end {
			arr[jmpIndex[i]].instruction = "jmp"
		} else {
			fmt.Printf("Parttwo: Reached end:	%v   with value:	%d", end, value)
			return
		}
	}
	for i := 0; i < len(nopIndex)-1; i++ {
		arr[jmpIndex[i]].instruction = "jmp"
		value, end := instructionHandler(arr)
		if !end {
			arr[jmpIndex[i]].instruction = "nop"
		} else {
			fmt.Printf("Parttwo: Reached end:	%v   with value:	%d\n", end, value)
			return
		}
	}

}

func main() {
	partone("input.txt")
	parttwo("input.txt")

}
