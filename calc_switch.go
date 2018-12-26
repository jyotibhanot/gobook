// Program to perform arithmetical operations. 
// For now we will perform four operations: addition, subtraction, division,
// multiplication.

package main
import (
	"fmt"
        "os"
	"strconv"
)

func main() {
	op1, _ := strconv.Atoi(os.Args[1])
	op2, _ := strconv.Atoi(os.Args[2])
        operator := string(os.Args[3])
	result := 0;
	
	switch operator {
		case "*": 
			result = mul(op1, op2)
		case "/":
			result = div(op1, op2)
		case "+":
			result = sum(op1, op2)
		case "-":
			result = diff(op1, op2)
		default: 
			fmt.Println("Wrong operator, you idiot!")
	}
	fmt.Println("Result: ", result)
}

func sum(a int, b int) int {
	return a+b
}

func diff(a int, b int) int{
	return a-b
}

func mul(a int, b int) int{
	return a*b
}

func div(a int, b int) int{
	return a/b
}
