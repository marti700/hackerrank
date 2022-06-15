package main

import "fmt"


func reverseNumber(n int) int {
	var rev int
	for n > 0  {
		rev = (rev*10) + n%10
		n = n/10
	}
	return rev
}

func isBeutiful(n,k int) bool {
	return (n-reverseNumber(n)) % k == 0
}

func beutifulDays(start, end, divisor int) int {
  var	beutifulDays int
	for i := start; i<=end; i++ {
		if isBeutiful(i,divisor) {
			beutifulDays++
		}
	}
	return beutifulDays
}


func main() {
	var start, end, divisor int

	fmt.Scanf("%d %d %d\n", &start, &end, &divisor)
	fmt.Println(beutifulDays(start, end, divisor))
}