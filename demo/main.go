package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	defer fmt.Println(1)
	defer fmt.Println(2)

	defer fmt.Println(3)

}

/* func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	for right != left {
for right>left&&
	}
}
*/
