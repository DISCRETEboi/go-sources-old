package main

import "fmt"
import "res1"

func main() {
	one := []string{"mat", "wale"}
	two := []string{"dayo", "mat"}
	test := res1.StrEqual(one, two)
	fmt.Println(test)
}