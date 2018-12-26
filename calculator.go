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
	
        if result := 0; operator == "+" {
		result = sum(op1, op2)
	} else if operator == "*" {
		result = mul(op1, op2)
	} else if operator == "-" { 
		result = diff(op1, op2)
	} else if operator == "/" {
		result = div(op1, op2)
	} else {
		fmt.Println("Wrong operator, you idiot!")
		result = 0
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
