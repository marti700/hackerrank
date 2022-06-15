package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func timeConversion(arr []string) string {
	if strings.TrimSpace(arr[2][2:]) == "PM" {
		if arr[0] != "12" {
			hour, _ := strconv.ParseInt(arr[0], 10, 0)
			arr[0] = strconv.Itoa((int(hour) + 12))
		}
	}

	if strings.TrimSpace(arr[2][2:]) == "AM" && arr[0] == "12" {
		arr[0] = "00"
	}

	return arr[0] + ":" + arr[1] + ":" + arr[2][0:2]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	time := strings.Split(input, ":")
	fmt.Print(timeConversion(time))
}
