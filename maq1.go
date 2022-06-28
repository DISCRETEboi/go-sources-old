package main

import "fmt"
import "res1"

func main() {
	//m1 := []int{}
	mat := []int{4,2,6,3,1,6,8,6,4,8,2,6,4,8,6}
	mat1 := []int{2,5,4,7,5,7,4}
	mat2 := []int{1,2}
	//mat3 := []int{}
	fmt.Println("+-------------------------------------------------+")
	fmt.Println(res1.Pop(mat1))
	fmt.Println(res1.Extend(mat1, mat2))
	fmt.Println(res1.Reverse(mat1))
	fmt.Println(res1.Count(mat, 2))
	fmt.Println(res1.RemoveAll(mat1, 4))
	fmt.Println(res1.Remove(mat1, 4, 1))
	fmt.Println(res1.Insert(mat2, 900, 1))
	var1 := res1.Copy(mat1)
	fmt.Println(var1, len(var1), cap(var1))
	fmt.Println(res1.Subset(mat1, []int{2,7,5}))
	fmt.Println(res1.Str2ints("adebowale"))
	var2 := res1.ToSet(mat1)
	fmt.Printf("%v , %[1]T\n" , var2)
	fmt.Println(res1.Member(var2, 15))
	fmt.Println(res1.LenSet(var2))
	fmt.Println("+-------------------------------------------------+")
}