// Implementing echo in go

package main

import "fmt"
import "os"

func main() {
	sep := ""
	var s string
 
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "	
	}
	fmt.Println(s)
}
