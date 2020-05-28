/* Use of pointers in Go */

package main

import "fmt"

func main(){

	i := 5
	inc(i)					// here the value of i is being passed by value and not reference!
	fmt.Println(i) 			//prints the value of i
	fmt.Println(&i)			//prints the memory address of i

	fmt.Println("Passing by Reference now!")
	ref_inc(&i)
	fmt.Println(i)
	fmt.Println(&i)

	/* the memory address will stay the same, since the variable at that specific memory address is being passed
	as reference. Only the value being stored inside it will change.
	 */
}

func inc(x int){
	x++
}

/*Though the function has no return type, the value is still incremented! This is the beauty of pointers, and passing
values by reference always helps in scenarios such as this.
 */

func ref_inc(x *int){
	*x++
}
