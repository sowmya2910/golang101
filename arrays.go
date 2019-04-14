package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	arr[4] = 75
	fmt.Println("set: ", arr)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	twoD := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D: ", twoD)
}
