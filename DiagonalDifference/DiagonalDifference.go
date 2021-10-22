package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func toIntArray(arr []string) []int64 {
  var r []int64

  for _, v := range arr {
    temp, _ := strconv.ParseInt(strings.TrimSpace(v), 0, 64)
    r = append(r, temp)
  }
  return r
}

//gets the left diagonal
func diagL(m [][]int64) int64 {
  var r, c, res int64
  mLen := len(m)

  for r < int64(mLen) && c < int64(mLen) {
    res += m[r][c]
    r++
    c++
  }
  return res
}

//gets the left diagonal
func diagR(m [][]int64) int64 {
  mLen := len(m)
  var r, res int64
  c := int64(mLen)-1

  for r <= int64(mLen) && c >= 0 {
    res += m[r][c]
    r++
    c--
  }
  return res
}

func abs(x int64) int64 {
  if x < 0 {
    return x * -1
  } else {
    return x
  }
}

func main() {
  var mDimension int
  matrix := [][]int64{}

  fmt.Scanf("%v\n", &mDimension)

  //builds the matrix
  for i := 0; i < mDimension; i++ {

    reader := bufio.NewReader(os.Stdin)
    row, _ := reader.ReadString('\n')
    matrix = append(matrix, toIntArray(strings.Split(row, " ")))
  }

  // fmt.Println(matrix)
  // fmt.Println(diagL(matrix))
  // fmt.Println(diagR(matrix))
  fmt.Println(abs(diagL(matrix) - diagR(matrix)))
}
