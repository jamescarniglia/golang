// basic switch or case type statements

package main

import "fmt"
import "time"

func main() {
	i := 2
	// no new line
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	// mulitple cases
	case time.Saturday, time.Sunday:
		fmt.Println("Oh my tiss the weekend!")
	default:
		fmt.Println("its a weekday")
	}

	// sorta like an if statment
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Hey I am a bool you fool!")
		case int:
			fmt.Println("Hey I am a int you fool!")
		default:
			// string substatution in a print statement
			fmt.Printf("really have no idea what you are, what is %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("sup")

}
