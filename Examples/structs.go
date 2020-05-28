/* Demonstration of structures in Go */

package main

import "fmt"

/* use the key word type when declaring the struct outside the scope of the function, followed by the name of the
structure and lastly use the key word struct. A structure can hold any type of variables inside it, including other
structures too. For now, we use only strings and int.
 */
type person struct {
	name string
	age int
}

func main(){

	p := person{name: "John", age: 25} 		//declare a struct p of type person. Short-hand declaration here, again.

	fmt.Println(p)
	fmt.Println(p.name)				//only print the name inside it and not anything else.
	fmt.Println(p.age)
}