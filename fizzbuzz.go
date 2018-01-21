package main

import "fmt"

func main() {

	// switch1
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		fmt.Print(" ")
	}

	//// switch2
	//for i := 1; i <= 100; i++ {
	//	switch i % 15 {
	//	case 0:
	//		fmt.Print("FizzBuzz")
	//	case 3, 6, 9, 12:
	//		fmt.Print("Fizz")
	//	case 5, 10:
	//		fmt.Print("Buzz")
	//	default:
	//		fmt.Print(i)
	//	}
	//	fmt.Print(" ")
	//}

	//// if
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 && i%5 == 0 {
	//		fmt.Print("FizzBuzz")
	//	} else if i%3 == 0 {
	//		fmt.Print("Fizz")
	//	} else if i%5 == 0 {
	//		fmt.Print("Buzz")
	//	} else {
	//		fmt.Print(i)
	//	}
	//	fmt.Print(" ")
	//}
}
