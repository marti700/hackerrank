package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func angryProfessor(k int, a []int) string {
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			k--
		}
		if k <= 0 {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	var t, n, k int // test cases, number of students, the threshold
	var r []string

	fmt.Scanf("%d\n", &t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d\n", &n, &k)
		reader := bufio.NewReader(os.Stdin)
		arrivalTimes, _ := reader.ReadString('\n')
		arrivalTimesArr := strings.Split(arrivalTimes, " ")
		var a = make([]int, n)
		for j := 0; j < len(arrivalTimesArr); j++ {
			temp, _ := strconv.Atoi(strings.TrimSpace(arrivalTimesArr[j]))
			a[j] = temp
		}
		r = append(r, angryProfessor(k, a))
	}

	for i, e := range r {
		if i == len(r)-1 {
			fmt.Print(e)
		} else {
			fmt.Println(e)
		}
	}
}
