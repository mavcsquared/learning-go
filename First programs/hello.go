package main

import (
	"fmt"
)

func main() {

	var myarray [250]int
	size := 0

	for x := 0; x <= 100; x++ {
		if x%2 != 0 {
			myarray[size] = x
			size++
		}
	}

	fmt.Println(myarray[0:size])

	anotherOne := [3]int{1, 2, 3}

	fmt.Println(anotherOne[1:2])

}
