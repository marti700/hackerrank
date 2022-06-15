package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func sum(arr []int) int {
	var sum int
	for _, e := range arr {
		sum += e
	}
	return sum
}

func sharedFairly (notSharedItem, payedByAnna int, bill []int) bool {
	billTotal := sum(bill)

	if ((billTotal - bill[notSharedItem]) / 2) == payedByAnna {
		return true
	}

	return false
}

// IO
func parseInput() (int, int, []int ) {
	reader := bufio.NewReader(os.Stdin)

	notShared_str, _ := reader.ReadString('\n')
	notShared_str = strings.TrimSpace(notShared_str)
	notShared_str_arr := strings.Split(notShared_str, " ")
	notShared, _ := strconv.Atoi(notShared_str_arr[1]) // the item Anna didn't share

	// bill arr

	bill_str, _ := reader.ReadString('\n')
	bill_str_arr := strings.Split(bill_str, " ")
	var bill []int
	for _, e := range bill_str_arr {
		temp, _ := strconv.Atoi(strings.TrimSpace(e))
		bill = append(bill, temp)
	}

	// the money Brian charge Anna
	annaPayed_str, _ := reader.ReadString('\n')
	annaPayed, _ := strconv.Atoi(strings.TrimSpace(annaPayed_str))


	return notShared, annaPayed, bill

}
 // IO END


func main () {
	notSharedItem, payedByAnna, bill := parseInput()
	fairlySplit := sharedFairly(notSharedItem, payedByAnna, bill)

	if fairlySplit{
		fmt.Println("Bon Appetit")
	} else {
		// surcharge = payedByAnna - [(billTotal - bill[notShareditem])/2]
		fmt.Println(payedByAnna - ( (sum(bill) - bill[notSharedItem])/2 ) )
	}
}