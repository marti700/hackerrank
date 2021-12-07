package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solves the problem by reversing the player array and searching through the ranked array to determine
// how the player did each time he/she played the game

// Solves the problem by reversing the player array and searching through the ranked array to determine
// how the player did each time he/she played the game

// E.X
// 100 100 50 40 40 20 10 --> ranked
// 5 25 50 120 --> player

// reverse player
// 120 50 25 5 --> player

//loop throgh player score loking how player did

// [120] 100 100 50 [50] 40 40 [25] 20 10 [5]

func climbingLeaderBoard(ranked, player []int) []int {
	var currentStandScore, currentStand int
	var playerPlaces []int
	var cachedIndex int // used to not traverse the ranked array every time

	// traverse the player scores slice from the back to the front so we start serching for the player's highest standing
	// in the leader board
	for i := len(player) - 1; i > -1; i-- {
		playerCurrentScore := player[i]
		for j := cachedIndex; j < len(ranked); j++ {
			// if the current score is diferent from the one being evaluated it means the score have change and with it the ranking
			if currentStandScore != ranked[j] {
				currentStandScore = ranked[j]
				currentStand++
			}
			if playerCurrentScore >= currentStandScore {
				playerPlaces = append(playerPlaces, currentStand)
				break
			}
			cachedIndex++
		}
	}
	// completes the playerPlacers slice with the not procced results
	// when the highest ranked score is higher than a particular player score
	// it wont be evaluated by the loop above.
	// from this point on all player scores are less than all the scores in the leader board
	if len(playerPlaces) < len(player) {
			currentStand++
		for ; len(playerPlaces) != len(player); {
			playerPlaces = append(playerPlaces, currentStand)
		}
	}

	return playerPlaces
}

// func placePlayer(ranked, player []int) int {

// }

func abs(x int) int {
	if x < 0 {
		return x * (-1)
	}
	return x
}

// IO

//Reads space separated integers from stdin
func getInput(reader *bufio.Reader) []int {
	arr_string, _ := reader.ReadString('\n')
	arr_split := strings.Split(arr_string, " ")
	var arr []int

	for _, e := range arr_split {
		temp, _ := strconv.Atoi(strings.TrimSpace(e))
		arr = append(arr, temp)
	}

	return arr
}

// prints results to sdout
func output(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}

// IO END

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') // ignores first line of input
	leaderBoard := getInput(reader)
	reader.ReadString('\n') // ignores third line of input
	playerScore := getInput(reader)
	// prints resutl to stdout

	output(climbingLeaderBoard(leaderBoard, playerScore))
}
