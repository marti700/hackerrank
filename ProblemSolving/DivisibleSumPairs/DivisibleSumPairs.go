package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func divisibleSumPairs(k int64, arr []int64) int64 {
	var pairs int64

	for i, _ := range arr {
		for j, _ := range arr {
			if (i < j) && ((arr[i] + arr[j]) % k == 0) {
				pairs++
			}
		}
	}

	return pairs
}

// IO HANDLING
func parseInput() (int64, []int64) {
	reader := bufio.NewReader(os.Stdin)
	// get the value of k
	n_k_in, _ := reader.ReadString('\n')
	n_k_str := strings.Split(n_k_in, " ")
	k_str := n_k_str[1]
	k, _ := strconv.ParseInt(strings.TrimSpace(k_str), 10, 0)

	//get the valueo of arr
	arr_str, _ := reader.ReadString('\n')
	arr_splt := strings.Split(arr_str, " ")

	//convert arr elements to int64
	var arr []int64
	for _, e := range arr_splt {
		temp, _ := strconv.ParseInt(strings.TrimSpace(e), 10, 0)
		arr = append(arr, temp)
	}

	return k, arr
}

// IO HANDLING END

func main() {
	fmt.Println(divisibleSumPairs(parseInput()))
}
