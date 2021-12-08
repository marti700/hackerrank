package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hurdleRace(k int, height []int) int {
	var tallestHurdle int

	for _, h := range height {
		if h > tallestHurdle {
			tallestHurdle = h
		}
	}

	if tallestHurdle-k < 0 {
		return 0
	} else {
		return tallestHurdle - k
	}
}

func getHurdlesHeight(reader *bufio.Reader) []int {
	arr_string, _ := reader.ReadString('\n')
	arr_split := strings.Split(arr_string, " ")
	var arr []int

	for _, e := range arr_split {
		temp, _ := strconv.Atoi(strings.TrimSpace(e))
		arr = append(arr, temp)
	}

	return arr
}

// prints results to sdout
func output(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	inString, _ := reader.ReadString('\n')
	splitedInString := strings.Split(inString, " ")
	maxJumpHeight_s := splitedInString[1]
	maxJumpHeight_i,_ := strconv.Atoi(strings.TrimSpace(maxJumpHeight_s))
	hurdlesHeight := getHurdlesHeight(reader)

	fmt.Println(hurdleRace(maxJumpHeight_i, hurdlesHeight))
}
