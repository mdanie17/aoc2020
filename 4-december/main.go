package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			lines = append(lines, line)
			line = ""
		}
	}
	lines = append(lines, line) // Get the last line of the file
	return lines

}

func stringParser(arr []string) []map[string]string {
	var passportList []map[string]string
	for _, v := range arr {
		passport := make(map[string]string)
		trimmed := strings.TrimSpace(v)
		fields := strings.Split(trimmed, " ")
		for _, v2 := range fields {
			fieldValue := strings.Split(v2, ":")
			passport[fieldValue[0]] = fieldValue[1]
		}
		passportList = append(passportList, passport)
	}
	return passportList
}

func validation(arr []map[string]string) []map[string]string {
	var counter int
	var validList []map[string]string
	var valid bool
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for i, _ := range arr {
		for _, v := range fields {
			valid = true
			_, ok := arr[i][v]
			if !ok {
				valid = false
				//fmt.Printf("%s was not in the string %s\n", v, arr[i])
				counter = counter + 1
				break
			}
		}
		if valid {
			validList = append(validList, arr[i])
		}
	}
	return validList
}

func validationPart2(arr []map[string]string) int {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var counter int
	var validated bool
	for i := range arr {
		validated = false
		for _, v := range fields {
			switch v {
			case "byr":
				number, _ := strconv.Atoi(arr[i][v])
				if (number < 1920) || (number > 2002) || (number < 999) {
					counter = counter + 1
					//fmt.Printf("invalid as 2002 < %d < 1920\n", number)
					validated = true
				}
			case "iyr":
				number, _ := strconv.Atoi(arr[i][v])
				if (number < 2010) || (number > 2020) || (number < 999) {
					counter = counter + 1
					//fmt.Printf("invalid as 2010 < %d < 2020\n", number)
					validated = true
				}
			case "eyr":
				number, _ := strconv.Atoi(arr[i][v])
				if (number < 2020) || (number > 2030) || (number < 999) {
					counter = counter + 1
					//fmt.Printf("invalid as  %d < 2020 || %d > 2030\n", number, number)
					validated = true
				}
			case "hgt":
				height := arr[i][v]
				if strings.Contains(height, "cm") {
					field := strings.Split(height, "cm")
					fieldInt, _ := strconv.Atoi(field[0])
					if fieldInt < 150 || fieldInt > 193 {
						counter = counter + 1
						//fmt.Printf("invalid as 150 < %d < 193\n", fieldInt)
						validated = true
					}
				} else if strings.Contains(height, "in") {
					field := strings.Split(height, "in")
					fieldInt, _ := strconv.Atoi(field[0])
					if fieldInt < 59 || fieldInt > 76 {
						counter = counter + 1
						//fmt.Printf("invalid as 59 < %d < 76\n", fieldInt)
						validated = true
					}
				} else {
					counter = counter + 1
					//fmt.Printf("Invalid as %s doesnt contain cm or in\n", height)
					validated = true
				}
			case "hcl":
				value := arr[i][v]
				//fmt.Println(value)
				_, err := strconv.ParseUint(value[1:], 16, 64)
				if err != nil || string(value[0]) != "#" {
					counter = counter + 1
					//fmt.Printf("invalid as %s is not hex\n", value)
					validated = true
				}
			case "ecl":
				value := arr[i][v]
				var isvalid bool
				valid := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				isvalid = false
				for _, v := range valid {
					if value == v {
						isvalid = true
					}
				}
				if !isvalid {
					counter = counter + 1
					//fmt.Printf("invalid as %s is not valid\n", value)
					validated = true
				}
			case "pid":
				value := arr[i][v]
				_, ok := strconv.Atoi(value)
				if !(len(value) == 9) || ok != nil {
					counter = counter + 1
					//fmt.Printf("Invalid as %s is not long enough\n", value)
					validated = true
				}
			case "cid":
			}
			if validated {
				break
			}
		}
	}
	return len(arr) - counter
}

func partOne(path string) {
	arr := fileReader(path)
	passports := stringParser(arr)
	answer := validation(passports)
	fmt.Println("First exercise: ", len(answer))
}

func partTwo(path string) {
	arr := fileReader("input.txt")
	passports := stringParser(arr)
	valid := validation(passports)
	answer := validationPart2(valid)
	fmt.Println("Second exercise: ", answer)
}

func main() {
	partOne("input.txt")
	partTwo("input.txt")
}
