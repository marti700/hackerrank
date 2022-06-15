package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func sumArray(ar[]string) int64{
  var r int64
  for _,v := range ar {
    temp, _ :=  strconv.ParseInt(strings.TrimSpace(v), 10, 0)
    r += temp
  }
  return r
}
func main (){
  var arrSize int

  fmt.Scanf("%v\n", &arrSize)
  reader := bufio.NewReader(os.Stdin)
  arrString, _ := reader.ReadString('\n')
  fmt.Println(sumArray(strings.Split(arrString, " ")))
}