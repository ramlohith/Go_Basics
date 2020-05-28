/* Program to set variables in Go */

package main

import(
	"fmt"
)

func main() {

	/* when uninitialized, the value is default to 0
	var x int */

	var x int = 5

	/* short hand for the same is shown below
 	go infers the values automatically */

	y := 7
	sum := x + y
	fmt.Println(sum)
}
