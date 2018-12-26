// Program to print day of the week

package main

import (
	"fmt"
	"os"
)

func main() {
	num := os.Args[1]
	day := ""
	if num == "1" {
		day = "Sunday"
	} else if num == "2" {
		day = "Monday!"
	} else if num == "3" {
		day =  "Tuesday"
	} else if num == "4" {
                day = "Wednesday"
        } else if num == "5" {
                day = "Thursday"
        } else if num == "6" {
                day = "Friday"
        } else if num == "7" {
                day = "Saturday"
        } else { 
    		fmt.Println("Wrong Input")
	}
	fmt.Println("Day is:", day)
}
