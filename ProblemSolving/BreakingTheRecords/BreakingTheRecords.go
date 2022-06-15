package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func breakingRecords (scores [] int64) [2] int64 {
	var maxScores [2] int64
	max := scores[0]
	min := scores[0]

	for _,elem := range scores {
		if (elem > max) {
			max = elem
			maxScores[0] ++
		}
		if (elem < min) {
			min = elem
			maxScores[1] ++
		}
	}

	return maxScores
}

func main () {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	input, _ := reader.ReadString('\n')

	scoresArr := strings.Split(input, " ")

	var sc [] int64
	for _, ele := range scoresArr {
		temp, _ := strconv.ParseInt(strings.TrimSpace(ele), 10, 0)
		sc = append(sc, temp)
	}

	result := breakingRecords(sc)
	fmt.Print(strconv.Itoa(int(result[0])) + " " + strconv.Itoa(int(result[1])))
}