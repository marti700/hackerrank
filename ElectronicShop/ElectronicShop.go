package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)


func getMoneySpent(budget int, keybords, usbs []int) int {
	maxExpent := -1

	for _, k := range keybords {
		for _, u := range usbs {
			spent := k + u
			if spent <= budget && spent > maxExpent {
				maxExpent = spent
			}
		}
	}
	return maxExpent
}

// IO
func readArrAsInt(str string) []int {
    var arr []int

    for _,e := range strings.Split(str, " ") {
        temp,_ := strconv.Atoi(strings.TrimSpace(e))
        arr = append(arr, temp)
    }

    return arr
}
func parseInput() (int, []int, []int) {
    reader := bufio.NewReader(os.Stdin)
    b_k_m, _ := reader.ReadString('\n')
    budget, _ := strconv.Atoi(strings.TrimSpace(strings.Split(b_k_m, " ")[0]))


    strKeyborad, _ := reader.ReadString('\n')
    strUSB, _ := reader.ReadString('\n')

    keybords := readArrAsInt(strKeyborad)
    usbs := readArrAsInt(strUSB)

    return budget, keybords, usbs
}

// IO END

func main() {
    fmt.Println(getMoneySpent(parseInput()))
}