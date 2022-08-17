package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'cost' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY B as parameter.
 */

func cost(arr []int) int {
    dp := make([][]int, len(arr))
    for i := range dp {
        dp[i] = make([]int, 2)
    }

    dp[0][0] = 0
    dp[0][1] = 0

    for i := 1; i < len(arr); i++ {
        dp[i][0] = maxx(dp[i-1][0]+abs(arr[i]-arr[i-1]), dp[i-1][1]+abs(arr[i]-1))
        dp[i][1] = maxx(dp[i-1][0]+abs(1-arr[i-1]), 0)
    }

    return maxx(dp[len(arr)-1][0], dp[len(arr)-1][1])
}
func maxx(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return a * (-1)
    }

    return a
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int(nTemp)

        BTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var B []int

        for i := 0; i < int(n); i++ {
            BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
            checkError(err)
            BItem := int(BItemTemp)
            B = append(B, BItem)
        }

        result := cost(B)

        fmt.Fprintf(writer, "%d\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
