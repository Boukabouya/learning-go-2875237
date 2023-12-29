package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // where the data come from (standard input)
	fmt.Print("Entere text :")
	input, _ := reader.ReadString('\n') /*By using _ as the variable name for the second return value, you are indicating that
	you are not interested in handling or checking the error explicitly.
	It's a way of saying "I know there might be an error,
	but I'm choosing not to handle it in this part of the code."*/
	// ReadString function requires a single character.
	fmt.Println("your entered is :", input)

}
