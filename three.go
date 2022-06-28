package main

import "fmt"

func main() {
	for count:=1; count<=100; count++ {
		det3 := count % 3
		det5 := count % 5
		if (det3 == 0) && (det5 == 0) {
			fmt.Println("FizzBuzz")
		} else if det3 == 0 {
			fmt.Println("Fizz")
		} else if det5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(count)
		}
	}
}