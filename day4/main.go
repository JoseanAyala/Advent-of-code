package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type winnerMap = map[string]struct{}

func getFileScanner() *bufio.Scanner {
	pathPtr := flag.String("path", "./input.txt", "Path")
	flag.Parse()

	file, err := os.Open(*pathPtr)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}

func parseGame(rawHand string, callBack func(num string)) {
	for left := 0; left < len(rawHand); left += 3 {
		right := left + 3
		if right > len(rawHand) {
			right = len(rawHand)
		}
		callBack(rawHand[left:right])
	}
}

func getPoints(player *[]string, winning *winnerMap) float64 {
	var handValue float64

	for _, num := range *player {
		if _, exists := (*winning)[num]; exists {
			if handValue > 1 {
				handValue *= 2
			} else {
				handValue++
			}
		}
	}

	return handValue
}

func main() {
	file := getFileScanner()

	var totalPoints float64
	for file.Scan() {
		playerNumbers := make([]string, 0)
		winningNumbers := make(winnerMap)

		raw := file.Text()
		game := raw[strings.Index(raw, ":")+1:]
		numbers := strings.Split(game, "|")

		parseGame(numbers[0], func(num string) {
			playerNumbers = append(playerNumbers, num)
		})

		parseGame(numbers[1], func(num string) {
			winningNumbers[num] = struct{}{}
		})

		totalPoints += getPoints(&playerNumbers, &winningNumbers)
	}

	fmt.Println(totalPoints)
}
