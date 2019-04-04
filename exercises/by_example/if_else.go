// simple if else
package main

import "fmt"

func main() {
	//if else
	if 7%2 == 0 {
		fmt.Println("7 is wicket even")
	} else {
		fmt.Println("7 is wicket odd")
	}

	// simple if
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//if else if
	if num := 1000; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has mulitple digits")
	}
}
