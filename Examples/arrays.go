/* Program to show fixed and dynamic arrays in Go */

package main

import (
	"fmt"
)

func main() {
	/* uninitialized arrays will all have values set to 0 */

	var a[5] int
	a[2] = 7
	/* Note, arrays begin counting from zero, thus 2nd element is actually the third! */

	/* Short-hand syntax for arrays is given below */
	b := [5] int{5,4,3,2,1}

	/* Dynamic arrays are called slices in Go, where the number of elements in an array are not fixed.
	An example for the same is shown below */

	c := [] int{1,2,3,4,5}

	/* we can use the built-in append function, to append an element to the array (or slice). Technically, this
	does not add a new element, but returns another array. But Go does all of that in the backend and we don't have to
	worry about that. */

	c = append(c, 6)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
