package main

import (
	"fmt"
)

func countValeys(path string) int{
	// number of valleys a step counter down steps substract up steps sum
	var valleys,steps int

	for _,p := range path {
		if p == 'D' {
			steps--
		} else {
			steps++
		}

		//when steps is -1 after a descent there will alwasy be a valley
		if p == 'D' && steps == -1 {
			valleys++
		}
	}
	return valleys
}

func main() {
	var n int
	var path string
	fmt.Scanf("%d\n%s\n", &n, &path)
	fmt.Println(countValeys(path))

}