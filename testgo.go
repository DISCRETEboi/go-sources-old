package main

import "fmt"
import "time"

func main() {
	start := time.Now()
	for i:=0; i <= 10000000; i++ {
		for j:=0; j <= 1000; j++ {
			//Do
		}
	}
	end := time.Now()
	fmt.Println("It used" , end.Sub(start))
}