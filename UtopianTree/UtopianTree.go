package main

import (
	"fmt"
)

//Brute force
func utopianTree(n int) int {
	spring := true
	treeHeight := 1
	for i := 0; i < n; i++ {
		if spring {
			treeHeight = treeHeight * 2
			spring = false
		} else {
			treeHeight++
			spring = true
		}

	}
	return treeHeight
}

// a more mathematical approach
// the strategy here started by just taking into consideration the odd cycles, which are the ones
// that get multiplied by two. For example if we have the following list of cycles

// 0|1|2|3|4|5 | 6| 7|8  |9  |10| 11   ---> Cycles
// 1|2|3|6|7|14|15|30|31 |62 |63| 126  ---> Growth

// if we calcuate the powers of two for each odd cycle we will get
// 1| 3| 5|7 | 9 |  11   --> the cycle
// 1| 2| 3|4 | 5 |  6    --> n
// 2| 4| 8|16| 32|  64   --> the result of taking 2^n
// 0| 4| 6|14| 30|  62   --> the difference between the actual cycle value and the one obtained by doing 2^n (we are calling this K)

// if we look really closely we would be able to realize that the difference between 2^n and K is always 2
// this fact conduct us to the formula used in the ecuation below

// ((2^n) *2)-2  --> this can give us the growth of the trees with odd cycles
// all is left to do now is figure out the value of n
// when n is even the value of n is just n/2 (we have the same number of odd and even numbers in a range fron 1 to n)
// when n is odd we have to add 1 to (n/2) because in a range from 1 to n where n is odd we'll have an extra odd number

// also in summer the trees grow +1. Summer always corresponds to even cycles
// what this function do is that it calculates the growth up to the previous number (which is odd) using the above given formula
// and adds 1 to the result.

func utopianTree2(n int) int {
	// return int((math.pow(2,n)*2) -2)
	x := (n / 2)
	if n%2 == 0 {
		return ((pow(2, x) * 2) - 2) + 1

	} else {
		return (pow(2, x+1) * 2) - 2
	}
}

func pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * pow(base, exp-1)
}

func main() {
	var tc, cycles int
	fmt.Scanf("%d/n", &tc)

	for i := 0; i < tc; i++ {
		fmt.Scanf("%d\n", &cycles)
		// fmt.Println(pow(2, cycles))
		fmt.Println(utopianTree2(cycles))
	}
}
