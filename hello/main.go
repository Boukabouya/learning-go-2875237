package main

import (
	"fmt"
)

const aConst int = 24

//aConst is unexported because it starts with a lowercase letter, so it can only be accessed within the same package (main, in this case)

func main() {
	var aString string = "Hello from Go !"
	fmt.Println(aString)
	fmt.Printf("The type of this variable is %T\n ", aString)
	var anInteger int = 9
	fmt.Println(anInteger) //explicitly declaration
	fmt.Printf("The type of this variable is %T\n", anInteger)
	var defautInteger int
	fmt.Println(defautInteger)
	fmt.Printf("The type of this variable is %T\n", defautInteger)
	var anotherString = "it's me Rayene" // implicitly declaration
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n", anotherString)
	var anotherInteger = 18
	fmt.Println(anotherInteger)
	fmt.Printf("The variable's type is %T\n", anotherInteger)
	// colon equals operator :
	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n", myString)
	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T\n", aConst)
}
