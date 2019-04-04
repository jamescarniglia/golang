// showing simple for loops in go
package main

import "fmt"

func main() {
	i := 1
	// go doesn't have a while loop, so
	// this is what we have
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// set value, test value, add value
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// breaking
	for {
		fmt.Println("HELLO")
		break
	}

	// how to continue to next iteration of loop
	for hi := 0; hi <= 5; hi++ {
		// if divisable by 2 aka even number then skip
		if hi%2 == 0 {
			continue
		}
		fmt.Println(hi)
	}
}
