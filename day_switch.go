// Program to print day of the week

package main

import (
	"fmt"
	"os"
)

func main() {
	num := os.Args[1]
	day := ""
	switch num {
		       case "1": 
				day = "Sunday"
                       case "2":
                                day = "Monday"
                       case "3":
                                day = "Tuesday"
                       case "4":
                                day = "Wednesday"
                       case "5":
                                day = "Thursday"
                       case "6":
                                day = "Friday"
                       case "7":
                                day = "Saturday"
                       default:
                                day = "Wrong day"
	}
	fmt.Println("Day is:", day)
}
