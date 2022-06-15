package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func migratoryBirds(arr []int) int {
	// keeps the frecuency table
	// the indices of this slice represents the id of the birds
	// and its values the number of times it was seen
	var sightings = make([]int, 200003)

	var mostFrecuent int //most frecuent brid id
	var frecuency int    // keeps track of the actual frecuency of the most sighted bird

	for _, e := range arr {
		sightings[e]++
		if sightings[e] >= frecuency {
			// if the current observation is equal to a previous one with a lesser id
			if (e > mostFrecuent) && (sightings[e] == frecuency) {
				// jsut continiue
				continue
			}
			frecuency = sightings[e]
			mostFrecuent = e
		}
	}
	return mostFrecuent
}

// IO
func parseInput() []int {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // ignore first line of input

	birds_str, _ := reader.ReadString('\n')
	birds_temp := strings.Split(birds_str, " ")

	var birds []int
	for _, e := range birds_temp {
		temp, _ := strconv.Atoi(strings.TrimSpace(e))
		birds = append(birds, temp)
	}

	return birds
}

// IO END

func main() {
	fmt.Println(migratoryBirds(parseInput()))
}
