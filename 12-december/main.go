package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type waypoint struct {
	xcoordinate int
	ycoordinate int
}

func (w *waypoint) verticalMove(amount int) {
	w.ycoordinate += amount
}

func (w *waypoint) horizontalMove(amount int) {
	w.xcoordinate += amount
}

func (w *waypoint) rotateR(amount int) {
	switch amount {
	case 90:
		tempx := w.xcoordinate
		w.xcoordinate = w.ycoordinate
		w.ycoordinate = -tempx
	case 180:
		w.xcoordinate = -w.xcoordinate
		w.ycoordinate = -w.ycoordinate
	case 270:
		tempx := w.xcoordinate
		w.xcoordinate = -w.ycoordinate
		w.ycoordinate = tempx
	}
}

func (w *waypoint) rotateL(amount int) {
	w.rotateR(360 - amount)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func fileReader(path string) []string {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func shipMover(arr []string) int {
	var xposition, yposition, heading int
	for _, v := range arr {
		value, err := strconv.Atoi(v[1:])

		if err != nil {
			log.Fatal(err)
		}
		switch v[0] {
		case 'N':
			yposition += value
		case 'S':
			yposition -= value
		case 'E':
			xposition += value
		case 'W':
			xposition -= value
		case 'R':
			heading += value
		case 'L':
			heading -= value
			if heading < 0 {
				heading += 360
			}

		case 'F':
			if (heading/90)%4 == 0 {
				heading = 0
				xposition += value
			} else if heading == 90 || (heading/90)%4 == 1 {
				heading = 90
				yposition -= value
			} else if (heading/90)%2 == 0 {
				heading = 180
				xposition -= value
			} else if heading/90%3 == 0 {
				heading = 270
				yposition += value
			}
		}
	}
	return abs(xposition) + abs(yposition)
}

func partone(path string) {
	arr := fileReader(path)
	fmt.Println("Part 1: ", shipMover(arr))
}

func parttwo(path string) {
	arr := fileReader(path)
	waypoint := waypoint{
		xcoordinate: 10,
		ycoordinate: 1,
	}
	var xposition, yposition int
	for _, v := range arr {
		value, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch v[0] {
		case 'N':
			waypoint.verticalMove(value)
		case 'S':
			waypoint.verticalMove(-value)
		case 'E':
			waypoint.horizontalMove(value)
		case 'W':
			waypoint.horizontalMove(-value)
		case 'F':
			xposition += value * waypoint.xcoordinate
			yposition += value * waypoint.ycoordinate
		case 'R':
			waypoint.rotateR(value)

		case 'L':
			waypoint.rotateL(value)
		}
	}
	fmt.Println("Part 2: ", abs(xposition)+abs(yposition))
}

func main() {
	partone("input.txt")
	parttwo("input.txt")
}
