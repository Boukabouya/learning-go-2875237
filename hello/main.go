package main

import (
	"fmt"
)

func main() {
	fmt.Println("Create loops with for statements")
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	// LOOP WITH INDEX
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	// LOOP WITH RANGE ( dont care about the index , only element)
	for i := range colors {
		fmt.Println(colors[i])
	}
	/* this loop is equivalent to the second loop in your previous question
	   but with the explicit indication that the index is not being used
	   (_ is a valid identifier, and it's commonly used as a throwaway variable name
	    when the value is not needed).*/
	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value :", value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum :", sum)
		if sum > 200 {
			goto theEnd // break or continue

		}
	}
theEnd:
	fmt.Println("End of the program")
}
