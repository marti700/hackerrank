package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func printspaces(spaces int) {
  for i := 0; i<spaces;i++ {
    print("----")
    fmt.Print(" ")
  }
}

func printhastag(hastags int) {
  var tags string
  for i := 0; i<hastags;i++ {
    tags += "#"
  }
  fmt.Println(tags)
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  stairCaseSize, _ := strconv.ParseInt(strings.TrimSpace(input),10,0)
  for i := 1; i<=int(stairCaseSize);i++ {
    printspaces(int(stairCaseSize)-i)
    printhastag(i)
  }
}