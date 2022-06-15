package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(xs []int64) int64 {
	var sum int64
	for _, e := range xs {
		sum += e
	}
	return sum
}
func chocolates(bar []int64, d, m int64) int64 {
	var chocolateSqr int64
	for i, _ := range bar {
		upperIndex := i + int(m)

		// prevents indexout of range error
		if upperIndex > len(bar) {
			break
		}

		if sum(bar[i:upperIndex]) == d {
			chocolateSqr++
		}
	}
	return chocolateSqr
}

// IO Handling
func stringSliceToIntSlice(s []string) []int64 {
	var slice []int64
	for _, e := range s {
		temp, _ := strconv.ParseInt(strings.TrimSpace(e), 10, 0)
		slice = append(slice, temp)
	}
	return slice
}
func parseInput() ([]int64, int64, int64) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	strconv.ParseInt(strings.TrimSpace(input), 10, 0) // chocolate bar squeres. Not needed

	var chocolate []int64

	//parse chocolate
	chocolateString, _ := reader.ReadString('\n')
	chocolateTemp := strings.Split(chocolateString, " ")
	chocolate = stringSliceToIntSlice(chocolateTemp)
	//parse birth day and birth month
	d_m, _ := reader.ReadString('\n')
	d_mTemp := strings.Split(d_m, " ")
	d, _ := strconv.ParseInt(strings.TrimSpace(d_mTemp[0]), 10, 0)
	m, _ := strconv.ParseInt(strings.TrimSpace(d_mTemp[1]), 10, 0)

	return chocolate, d, m
}

// IO END

func main() {
	fmt.Println(chocolates(parseInput()))
}
