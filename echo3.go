// Using range

package main

import "fmt"
import "os"

func main() {
	s, sep := "", ""
   	// range() returns 2 values -- (index, value) in case of array, list  --- (key, value) in case of dict, map.
	for _, j := range(os.Args[1:]) {
		s += sep + j
		sep = " "
	}
	fmt.Println(s)
}
