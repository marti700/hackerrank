package main

import (
	"fmt"
)

func abs (x int) int {
	if x < 0 {
		return x * (-1)
	}

	return x
}


func catAndMouse(x,y,z int) string {
	c1 := abs(x-z)
	c2 := abs(y-z)

	if c1 < c2 {
		return "Cat A"
	} else if c2 < c1 {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func main() {
	var q int

	fmt.Scanf("%d\n", &q);

	for i := 0; i< q; i++{
		var x,y,z int
		fmt.Scanf("%d %d %d\n", &x, &y, &z )
		fmt.Println(catAndMouse(x,y,z))
	}
}