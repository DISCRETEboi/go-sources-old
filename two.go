package main

import "fmt"

func main() {
	for count:=1; count<=100; count++ {
		if count % 3 == 0 {
			fmt.Println("The number," , count , "is divisible by 3")
		}
	}
}