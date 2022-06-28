package main

import "fmt"
import "sort"
import "time"

//This is my implementation of the bubble sort algorithm
//And it's execution time is even shorter than that of the builtin 'sort.Ints()'
func main() {
	arr := [10]int{}
	i := 0
	for i < 10 {
		fmt.Print("Enter the value at index " , i+1 , " ---> ")
		fmt.Scanf("%d" , &arr[i])
		i++
	}
	fmt.Println("The array to sort is" , arr)
	arr1 := arr
	start := time.Now() //This is to time-test
	bsort(arr[:])
	end := time.Now() //This is to time-test
	t := end.Sub(start) //This is to time-test
	start1 := time.Now() //This is to time-test
	sort.Ints(arr1[:])
	end1 := time.Now() //This is to time-test
	t1 := end1.Sub(start1) //This is to time-test
	fmt.Println("My implementation sorted the array as" , arr , "within" , t)
	fmt.Println("Go's implementation sorted the array as" , arr1 , "within" , t1)
}

func bsort(arg []int) {
	for {
		n := 0
		for a, b := 0, 1; b <= len(arg)-1; a, b = a+1, b+1 {
			if arg[a] > arg[b] {
				arg[a], arg[b] = arg[b], arg[a]
				n++
			}
		}
		if n == 0 {
			return
		}
	}
}