package main

import "fmt"

func FizzBuzz() {
	i := 1
	for i <= 100 {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Println(i, ": FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, ": Fizz")
		case i%5 == 0:
			fmt.Println(i, ": Buzz")
		}
		i++
	}
}
