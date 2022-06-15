package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sockMerchatn(socks []int) int {
	var s = make([]int, 102)
	var pairs int

	for _, e := range socks {
		s[e]++
		if s[e]%2 == 0 {
			pairs += s[e] / 2
			s[e] = 0
		}
	}
	return pairs
}

// IO
func parseInput() []int {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // ignores first line of input

	socks_as_str, _ := reader.ReadString('\n')
	socks_arr := strings.Split(socks_as_str, " ")
	var socks []int

	for _, e := range socks_arr {
		temp, _ := strconv.Atoi(strings.TrimSpace(e))
		socks = append(socks, temp)
	}

	return socks
}

// IO END

func main() {
	fmt.Println(sockMerchatn(parseInput()))
}
