package main

import (
	"fmt"
)

// Simple fizzbuzz implementation
func main() {

	//print number from 1 - 100
	for i := 1; i < 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			//divisible by both three and five , print “FizzBuzz"

			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			// divisible by 3 print "Fizz"
			fmt.Println("fizz")

		} else if i%5 == 0 {
			// divisible by 5 print "buzz"
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
