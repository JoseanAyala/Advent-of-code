package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func scanFile() *bufio.Scanner {
	pathPtr := flag.String("path", "./input.txt", "Path")
	flag.Parse()

	file, err := os.Open(*pathPtr)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

func readFile(scanner *bufio.Scanner) *[]string {
	coordinates := make([]string, 0)

	for scanner.Scan() {
		puzzle := scanner.Text()
		var first, last string

		for _, c := range puzzle {
			char := string(c)
			_, err := strconv.Atoi(char)

			if err == nil {
				if len(first) == 0 {
					first = char
				}
				last = char
			}
		}

		coordinates = append(coordinates, first+last)
	}

	return &coordinates
}

func main() {
	coordinates := readFile(scanFile())

	finalSum := 0
	for _, coor := range *coordinates {
		coorNum, err := strconv.Atoi(coor)
		if err == nil {
			finalSum += coorNum
		}
	}

	fmt.Print("Final ammount is " + strconv.Itoa(finalSum))
}
