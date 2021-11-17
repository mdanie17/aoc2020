package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fileReader(path string) []string {
	var line string
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			line += fmt.Sprintf("%s ", scanner.Text())
		} else {
			lines = append(lines, strings.TrimSpace(line))
			line = ""
		}
	}
	lines = append(lines, strings.TrimSpace(line)) // Get the last line of the file
	return lines

}

func answerCounter(answers string) int {
	answerList := make(map[string]bool)
	for _, v := range answers {
		if string(v) != " " {
			answerList[string(v)] = true
		}
	}
	return len(answerList)
}

func newanswerCounter(answers string) int {
	var personcounter int = 1
	answerList := make(map[string]int)
	var counter int
	for _, v := range answers {
		if string(v) != " " {
			answerList[string(v)] += 1
		} else {
			personcounter += 1
		}
	}
	for i, _ := range answerList {
		if answerList[i] == personcounter {
			counter += 1
		}
	}
	return counter
}

func partone(path string) {
	var counter int
	lines := fileReader(path)
	for _, v := range lines {
		counter += answerCounter(v)
	}
	fmt.Println(counter)
}

func parttwo(path string) {
	var counter int
	lines := fileReader(path)
	for _, v := range lines {
		counter += newanswerCounter(v)
	}
	fmt.Println(counter)
}

func main() {
	partone("input.txt")
	parttwo("input.txt")
}
