package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "bufio"
)

func sum(arr []int64) int64 {
  var res int64 = 0
  for _,v := range arr {
    res += v
  }
  return res
}

func toIntArray(arr[]string) []int64 {
  var r []int64

  for _,v :=  range arr {
    temp, _:= strconv.ParseInt(strings.TrimSpace(v),0,64)
    r = append(r,temp)
  }
  return r
}

func main(){
  var intN int64
  fmt.Scanf("%v\n", &intN)

  reader := bufio.NewReader(os.Stdin)
  arrInt, _ := reader.ReadString('\n')
  numElements := toIntArray(strings.Split(arrInt, " "))
  fmt.Print(sum(numElements))
}