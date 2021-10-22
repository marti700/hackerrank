package main

import (
	"fmt"
)

func solution (pos1, jump1, pos2, jump2 int) string {
	if (pos2 - pos1) / (jump1 - jump2) > 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main () {
	var sPos1, jump1, sPos2, jump2 int
	fmt.Scanf("%v %v %v %v\n", &sPos1, &jump1, &sPos2, &jump2)
	fmt.Println(sPos1, jump1, sPos2, jump2)
}