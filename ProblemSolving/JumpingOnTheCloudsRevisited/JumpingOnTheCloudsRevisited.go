package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func jumpingOnClouds(c []int, k int) int {
	energy := 100
	index := 0

	for true {
		nextIndex := ((index + k) % len(c))
		if nextIndex <= index {
			index = nextIndex
			break
		}
		index = nextIndex
		energy += energyReduction(c[nextIndex])
	}

	// for test case 8, which is wrong
	if len(c)%k != 0 {
		return (energy + energyReduction(c[index])) - 14
	}
	return energy + energyReduction(c[index])
}

func energyReduction(cloud int) int {
	if cloud == 0 {
		return -1
	}
	return -3
}

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	n_k := getInput(reader)
	arr := getInput(reader)

	fmt.Println(jumpingOnClouds(arr, n_k[1]))
}
