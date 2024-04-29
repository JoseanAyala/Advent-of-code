package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rgbMap = map[string]int

type invalidGames map[int]int

var maxRGB = rgbMap{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func scanFile() *bufio.Scanner {
	pathPtr := flag.String("path", "./input.txt", "Path")
	flag.Parse()

	file, err := os.Open(*pathPtr)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}

func readFile(scanner *bufio.Scanner) *[]int {
	possibleGames := make([]int, 0)
	gameIndex := 1

	for scanner.Scan() {
		roundMax := rgbMap{
			"red":    0,
			"green:": 0,
			"blue":   0,
		}
		isBadGame := false

		input := scanner.Text()
		inputs := strings.Split(input, ":")
		gameRemoved := inputs[1:]

		for _, round := range gameRemoved {
			rounds := strings.Split(round, ";")

			for _, roundResult := range rounds {
				numAndColor := strings.Split(roundResult, ",")

				for _, splitByPick := range numAndColor {
					trimmedPick := splitByPick[1:]
					pick := strings.Split(trimmedPick, " ")

					num, _ := strconv.Atoi(pick[0])
					color := pick[1]
					mapNum := roundMax[color]

					if num > mapNum {
						roundMax[color] = num
					}
				}
			}
		}

		for key, maxVal := range maxRGB {
			rVal, ok := roundMax[key]

			if ok && rVal > maxVal {
				isBadGame = true
			}

		}

		if !isBadGame {
			possibleGames = append(possibleGames, gameIndex)
		}

		gameIndex++
	}

	return &possibleGames
}

func main() {
	results := readFile(scanFile())

	var totalSum int
	for _, gameIndex := range *results {
		totalSum += gameIndex
	}

	fmt.Println("Result is", totalSum)
}
