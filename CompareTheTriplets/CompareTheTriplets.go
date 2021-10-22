package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "bufio"
)

func getPoints(a,b []string) [2]int {
  p := [2]int{0,0}
  for k, _ := range a {
    alicePoints, _ := strconv.ParseInt(strings.TrimSpace(a[k]), 10,0)
    bobPoints, _ := strconv.ParseInt(strings.TrimSpace(b[k]),10,0)
    if alicePoints > bobPoints {
      p[0] += 1
    } else if alicePoints == bobPoints{
      continue
    } else {
      p[1] += 1
    }
  }
  return p
}

func main () {
  readerAlice := bufio.NewReader(os.Stdin)
  readerBob := bufio.NewReader(os.Stdin)

  alicePoints, _  := readerAlice.ReadString('\n')
  bobPoints, _:= readerBob.ReadString('\n')

  r := getPoints(strings.Split(alicePoints, " "), strings.Split(bobPoints, " "))

  fmt.Print(strconv.Itoa(r[0]) + " " + strconv.Itoa(r[1]))
}