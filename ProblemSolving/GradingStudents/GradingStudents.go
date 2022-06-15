package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int64) int64{
	if x < 0 {
		return x * (-1)
	} else {
		return x
	}
}
func nextFiveMultiple(grade int64) int64 {
	gradelastDigit := grade% 10

	if gradelastDigit < 5 {
		return grade + abs((5 - (grade % 10)))
	} else if gradelastDigit == 5 {
		return grade + 5
	} else {
		return grade + (10 - (grade % 10))
	}
}

func round(grade int64) int64 {
	if grade < 38 {
		return grade
	} else if (abs(grade-nextFiveMultiple(grade))) < 3 {
		return nextFiveMultiple(grade)
	} else {
		return grade
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	gradesNum, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 0)

	var grades []int64
	for i := 0; i < int(gradesNum); i++ {
		grads, _ := reader.ReadString('\n')
		gradsAsInt, _ := strconv.ParseInt(strings.TrimSpace(grads), 10, 0)
		grades = append(grades, gradsAsInt)
	}

	for _, e := range grades {
		fmt.Println(round(e))
	}
}
