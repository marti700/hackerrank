package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// One can clearly realize that when a applying a number of rotations k to an array of lengh l
// at one point the array elements will be back at their starting position
// E.X.:
// After applying 3 rotations to a length 3 array :
// Original -> 1 2 3
// k = 1
// 3 1 2
// k = 2
// 2 3 1
// k = 3
// 1 2 3

// at rotation 3 the array elements are back to ther original positions, given the challenge constraints
// This problem is about finding the number of rotations left after the array elements are back to ther original positions

// this number is given by taking the modulo of the number of queries and the lenght of the array (k % lenght_of_array)
// once we got that number all we have to do is count backwards on the original array to get the correct answer

func circularArrayRotation(n []int, k int, queries []int) []int {
	results := make([]int, len(queries))
	for i, q := range queries {
		results[i] = get(n, k, q)
	}

	return results
}

func get(n []int, k, q int) int {
	arr_len := len(n)
	m := k % arr_len // after k rotations there are m rotions left to be applied
	diff := q - m    //count backwrds
	if diff < 0 {
		// start counting from array last element if diff is less that 0 by doing
		// arr_len - abs(diff) to get into consideration the amount of steps taken to 0
		return n[arr_len-int(math.Abs(float64(diff)))]
	} else {
		return n[diff]
	}
}

// FOR SOME REASON THIS INPUT DOES NOT WORK IN HACKERRANK SO COPY THE CODE ABOVE AND REPLACE IT WITH
// THE circularArrayRotation FUNCTION IN HACKERRANK EDITOR TO COMPLETE THE CHALLENGE

// IO

//Reads space separated integers from stdin
func getInput(reader *bufio.Reader) []int {
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
	for _, e := range arr {
		fmt.Println(e)
	}
}

// IO END

func main() {
	reader := bufio.NewReader(os.Stdin)
	nkq := getInput(reader) // the number of elements in the integer array (n), the rotation count (k) and the number of queries (q)
	arr := getInput(reader) // reads the array
	queries := make([]int, nkq[2])

	for i := 0; i < nkq[2]; i++ {
		var q int
		fmt.Scanf("%d\n", &q) // read the query q
		queries[i] = q
	}

	output(circularArrayRotation(arr, nkq[1], queries))
}
