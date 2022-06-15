package main

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"os"
)

// determines the index of this char in the letter array
func hash(c rune) int {
	return int(c)-97
}

func designerPdfViewer(letters []int, word string) int {
	var maxHeight int
	for _, char := range word {
		if maxHeight < letters[hash(char)] {
			maxHeight = letters[hash(char)]
		}
	}
	return maxHeight*len(word)
}

func getArr(reader *bufio.Reader) []int {
	arr_string, _ := reader.ReadString('\n')
	arr_split := strings.Split(arr_string, " ")
	var arr []int

	for _, e := range arr_split {
		temp, _ := strconv.Atoi(strings.TrimSpace(e))
		arr = append(arr, temp)
	}

	return arr
}

func main () {
	reader :=bufio.NewReader(os.Stdin)
	letters := getArr(reader)
	word,_ := reader.ReadString('\n')

	fmt.Println(designerPdfViewer(letters, strings.TrimSpace(word)))
}