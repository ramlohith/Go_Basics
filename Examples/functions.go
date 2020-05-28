/* Go program to demonstrate a simple use of functions and how they are to be declared */

package main

import (
	"errors"
	"fmt"
	"math"
)

func main(){

	result := sum(2,3) 	// store the result in a variable

	sq, err := sqrt(1)		// store the result and the error returned from our sqrt function

	fmt.Println(result)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sq)
	}
}

/* use the func key word to declare a new function, followed by the name and the parameters to be passed.
If the function does have a return type, mention the return type after the parameters are shown below.
The function below takes two parameters and returns their sum to the calling function.
 */

func sum(x int, y int) int {
	return x + y
}

/* We can have more than one return type or values in go. To achieve that, mention them after the paramters in ()
as shown below. Here we also see a new package errors being imported. This is similar to working with exceptions in Java
where we use the try{}catch{} clause. Go does not have exceptions and using the errors package is one way of dealing
with exceptions.
 */

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0 , errors.New("square root undefined for negative numbers")
	} else {
		return math.Sqrt(x), nil
	}
}