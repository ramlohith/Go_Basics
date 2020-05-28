/* Simple Go program to demonstrate use of maps (Dictionary in Python/ map in C++/ hashmap in Java */

package main

import "fmt"

func main(){
	/* vertices is the name of the map being used, followed by a short-hand declaration.
	use the make key word to create a map memory allocation, followed by key word map to specify the use of maps
	make(map[ << the type of the key is mentioned here >> ] [ << the type of value is mentioned here >> ] */

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["blah"] = 12

	/* use of built-in delete function on a map is shown below */
	delete(vertices, "blah")

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
}