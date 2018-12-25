//Sum of all the numbers separated by space

package main

import (
	"fmt"
	"strconv"
	"os"
)

func main(){
   var sum int
   for i:=1; i < len(os.Args); i++ {
   	// Empty variable is used to hold ignored values.
   	num, _ := strconv.Atoi(os.Args[i]) 
   	// strconv.atoi is used to convert string-representation of integer to int.
   	sum += num
  }
  fmt.Printf("Sum = %d\n" , sum)
}
