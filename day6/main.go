package main

import (
  "fmt"
)

var (
 recordTimes = []int{40, 70, 98, 79}
 distance    = []int{215, 1051, 2147, 1005}
)

func main() {
  possibleCombination := make([]int, 0)
	for i, recordT := range recordTimes {
		possibleTimes := make([]int, 0)

		for pressT := 1; pressT < recordT; pressT++ {
			velocity := 0
			distanceToGo := distance[i]

			for raceT := 0; raceT < recordT; raceT++ {
				if raceT < pressT {
					velocity++
				} else {
					distanceToGo -= velocity
				}

				if distanceToGo < 0 {
					fmt.Println(raceT, pressT, velocity, recordT, distanceToGo)
					possibleTimes = append(possibleTimes, pressT)
					break
				}
			}

		}

		possibleCombination = append(possibleCombination, len(possibleTimes))
		fmt.Println(possibleTimes)
	}

	total := 0
	for _, val := range possibleCombination {
		if total == 0 {
			total += val
		} else {
			if val == 0 {
				total *= 1
			} else {
				total *= val
			}
		}
	}
	fmt.Println(total)
}
