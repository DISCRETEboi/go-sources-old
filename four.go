package main

import "fmt"

func main() {
	mat := "mateen"
	matp := &mat
	fmt.Println("mat has the value" , mat , "at address" , matp)
	mat = mat + " qamar"
	matq := &mat
	fmt.Println("mat has the value" , mat , "at address" , matq)
	isEq := matp == matq
	fmt.Println(isEq)
	/*var arr []int
	arr = []int{2,4}
	fmt.Println(len(arr))*/
}