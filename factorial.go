// Program to print factorial of a number
package main
import(
	"fmt"
	"os"
	"strconv"
)
func main(){
	var fact int=1
	n, _ := strconv.Atoi(os.Args[1])
	for j := n; j >= 1; j-- {
		fact=fact*j
	}
	fmt.Println("factorial=",fact)
}
