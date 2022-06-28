package main

import "fmt"

func main() {
	//m0 := 'y'
	//m01 := 'ʡ'
	m1 := "ǂtŲǂyʡĨʍΛdǂ"
	for _, me := range m1 {
		fmt.Printf("%T %[1]v\n", me)
	}
	fmt.Println("+------------------+")
	for i:=0; i < len(m1); i++ {
		fmt.Printf("+   %T %[1]v %v    +\n" , m1[i], i+1)
	}
}