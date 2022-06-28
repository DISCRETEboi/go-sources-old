package main

import "fmt"
import "time"
//import "runtime"

func main() {
	fmt.Println("The number, 1 is not prime")
	fmt.Println("The number, 2 is prime")
	start := time.Now()
	var nums int
	for primeNo:=3; primeNo<=100000; primeNo++ {
		//runtime.Gosched()
		divisor := 2
		index := 0
		for divisor < primeNo {
			//runtime.Gosched()
			if primeNo % divisor == 0 {
				index++
			}
			divisor++
		}
		if index == 0 {
			fmt.Println("The number," , primeNo , "is prime")
			nums++
		} else {
			fmt.Println("The number," , primeNo , "is not prime")
		}
	}
	fmt.Println("There are" , nums , "prime numbers between 1 and 100000")
	end := time.Now()
	fmt.Println(end.Sub(start))
}