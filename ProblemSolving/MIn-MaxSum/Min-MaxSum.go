package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// func calcMaximun(arr[] int64) int64 {

// }

func arrToInt(arr[] string) []int {
  var intarr[] int
  for _,v := range arr {
    temp,_ := strconv.ParseInt(strings.TrimSpace(v),10,0)
    intarr = append(intarr,int(temp))
  }
  return intarr
}

func minSum(arr[] int) int {
  var r int
  for i :=0; i<len(arr)-1; i++ {
      r += arr[i]
  }
  return r
}

func maxSum(arr[] int) int {
  var r int
  for i :=len(arr)-1; i>0; i-- {
      r += arr[i]
  }
  return r
}


func main() {
  reader := bufio.NewReader(os.Stdin)
  line, _ := reader.ReadString('\n')
  inputasIntArr := arrToInt(strings.Split(line, " "))
  sort.Ints(inputasIntArr)

  fmt.Print(strconv.Itoa(minSum(inputasIntArr)) + " " + strconv.Itoa(maxSum(inputasIntArr)))
}