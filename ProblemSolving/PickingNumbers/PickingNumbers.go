package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return x * (-1)
	}
	return x
}

func filter(arr []int) []int {
	m := make(map[int]int)
	var numberpick []int
	var valid bool
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if _, found := m[arr[j]]; found {
				continue
			}
			if abs(arr[i]-arr[j]) > 1 {
				valid = false
				m[arr[i]] = 1
				break
			}
			valid = true
		}
		if valid {
			numberpick = append(numberpick, arr[i])
		}
	}
	return numberpick
}

func pickingNumbers(arr []int) int {
	var temp []int
	var maxLenght int
	for i := 0; i < len(arr); i++ {
		temp = append(temp, arr[i])
		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i]-arr[j]) <= 1 {
				temp = append(temp, arr[j])
			}
		}
		filteredSlice := filter(temp)
		if maxLenght < len(filteredSlice) {
			maxLenght = len(filteredSlice)
		}

		temp = make([]int, 0)

	}
	return maxLenght
}

// IO
func getInput() []int {
	reader := bufio.NewReader(os.Stdin)
	//ignores first line of input
	reader.ReadString('\n')

	arr_string, _ := reader.ReadString('\n')
	arr_split := strings.Split(arr_string, " ")
	var arr []int

	for _, e := range arr_split {
		temp, _ := strconv.Atoi(strings.TrimSpace(e))
		arr = append(arr, temp)
	}

	return arr
}

// IO END
func main() {

	input := getInput()
	fmt.Println(pickingNumbers(input))
}
