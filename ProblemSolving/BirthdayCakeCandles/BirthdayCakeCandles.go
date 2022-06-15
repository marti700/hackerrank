package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCandles(arr[] int64) int64 {
	var frequency int64
	var tallestCandle int64

	for _,v := range arr {
		if(v > tallestCandle){
			tallestCandle = v
			frequency = 0
		}
		if (v == tallestCandle) {
			frequency++
		}
	}
	return frequency
}

func toIntSlice(arr[] string) []int64 {
	var intSlice [] int64
	for _,v := range arr {
		temp,_ := strconv.ParseInt(strings.TrimSpace(v), 10,0)
		intSlice = append(intSlice, temp)
	}
	return intSlice
}

func main () {
	var a int
  fmt.Scanf("%v\n", &a)
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	inputsAsIntSlice := toIntSlice(strings.Split(input, " "))
	fmt.Print(getCandles(inputsAsIntSlice))
}