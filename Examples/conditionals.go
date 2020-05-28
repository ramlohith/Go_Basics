/* Program to show if-else conditions in go */

package main

import (
	"fmt"
)

func main() {
	x := 7

	/* only note here is that there are not brackets around the conditions */

	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {
		fmt.Println("less than 2")
	} else {
		fmt.Println("executing else")
	}
}
