package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type (
	grid       = [][]string
	coordinate struct {
		y int
		x int
	}
)

func getFileScanner() *bufio.Scanner {
	pathPtr := flag.String("path", "./input.txt", "Path")
	flag.Parse()

	file, err := os.Open(*pathPtr)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}

func isNumber(char string) bool {
	isNumber, _ := regexp.MatchString("[0-9]", char)
	return isNumber
}

func isSymbol(char string) bool {
	return !isNumber(char) && char != "."
}

func fileToMatrix(scanner *bufio.Scanner) (*grid, *[]coordinate) {
	symbolsToVerify := make([]coordinate, 0)
	matrix := make(grid, 0)
	var y int

	for scanner.Scan() {
		row := make([]string, 0)

		for x, char := range scanner.Text() {
			character := string(char)

			if isSymbol(character) {
				symbolsToVerify = append(symbolsToVerify, coordinate{
					y: y,
					x: x,
				})
			}

			row = append(row, character)
		}

		matrix = append(matrix, row)
		y++
	}

	return &matrix, &symbolsToVerify
}

func fetchNumber(centerPoint coordinate, matrixPtr *grid, visitedPtr *map[coordinate]bool) int {
	matrix := (*matrixPtr)
	visited := (*visitedPtr)

	leftOffset, rightOffset := -1, 1
	shouldStopLeft, shouldStopRight := false, false
	for {
		left := centerPoint.x + leftOffset
		visited[coordinate{y: centerPoint.y, x: left}] = true

		if left == 0 {
			shouldStopLeft = true
		}

		if isNumber(matrix[centerPoint.y][left]) {
			if !shouldStopLeft {
				leftOffset--
			}
		} else {
			leftOffset++
			shouldStopLeft = true
		}

		right := centerPoint.x + rightOffset
		visited[coordinate{y: centerPoint.y, x: right}] = true

		if right == len(matrix[centerPoint.y])-1 {
			shouldStopRight = true
		}

		if isNumber(matrix[centerPoint.y][right]) {
			if !shouldStopRight {
				rightOffset++
			}
		} else {
			rightOffset--
			shouldStopRight = true
		}

		if shouldStopLeft && shouldStopRight {
			break
		}
	}

	left := centerPoint.x + leftOffset
	right := centerPoint.x + rightOffset

	strNum := strings.Join(matrix[centerPoint.y][left:right+1], "")
	num, err := strconv.Atoi(strNum)
	if err != nil {
		log.Fatal("found string in cast")
	}

	return num
}

func processSymbol(centerPoint coordinate, matrixPtr *grid) int {
	directions := []coordinate{
		// Top row
		{
			y: -1, x: -1,
		},
		{
			y: -1, x: 0,
		},
		{
			y: -1, x: 1,
		},
		// Middle row
		{
			y: 0, x: -1,
		},
		{
			y: 0, x: 1,
		},
		// Bottom row
		{
			y: 1, x: -1,
		},
		{
			y: 1, x: 0,
		},
		{
			y: 1, x: 1,
		},
	}

	amount := 0
	matrix := (*matrixPtr)
	visited := make(map[coordinate]bool, 0)
	for _, offSet := range directions {
		point := coordinate{y: centerPoint.y + offSet.y, x: centerPoint.x + offSet.x}
		if point.y < 0 || point.y > len(matrix) || point.x < 0 || point.x > len(matrix[0]) || visited[point] {
			continue
		}

		cell := matrix[point.y][point.x]

		if isNumber(cell) {
			amount += fetchNumber(point, matrixPtr, &visited)
		}

		visited[point] = true
	}
	return amount
}

func main() {
	matrix, symbolsToVerify := fileToMatrix(getFileScanner())

	amount := 0
	for _, coord := range *symbolsToVerify {
		amount += processSymbol(coord, matrix)
	}
	fmt.Println("Result is", amount)
}
