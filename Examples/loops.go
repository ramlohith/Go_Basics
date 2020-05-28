/* Demonstration of use of loops in Go */

package main

import "fmt"

func main(){

	/* Everything in go is a for loop!
	A simple for loop is shown below. Use the short-hand to declare a variable inside the loop
	 */
	for i:= 0; i < 5; i++ {
		fmt.Println(i)
	}

	/* The for loop can be used as a while loop too. The code below demonstrates a simple use */

	fmt.Println("Checking how a while loop works!")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	/* A for loop can also be used to show how we can iterate over elements in an array/ slice
	The loop gives us both the index and the element value at the index when iterating, which can be useful
	*/

	fmt.Println("Looping Over elements of a slice!")

	arr := []string{"a", "b", "c"}

	for index,value := range arr {
		fmt.Println("index:", index, "value:", value)
	}


	/* A for loop can also be used to show how we can iterate over elements in a map
	The loop gives us both the key and the element value at the index when iterating, which can be useful
	*/

	fmt.Println("Looping over elements of a map now!")

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("Key:", key, "value", value)
	}
}