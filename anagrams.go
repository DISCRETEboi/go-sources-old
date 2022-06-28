package main

import "fmt"
import "sort"

func main() {
	var first string
	var second string
	fmt.Print("Enter the first text for comparison--> ")
	fmt.Scan(&first)
	fmt.Print("Enter the second text for comparison--> ")
	fmt.Scan(&second)
	flen := len(first)
	slen := len(second)
	if flen != slen {
		fmt.Println("The strings are not anagrams! (unequal length and unequivalent letters)")
	} else {
		f := make([]int, 0)
		s := make([]int, 0)
		for _, mat := range first {
			f = append(f, int(mat))
		}
		for _, mat := range second {
			s = append(s, int(mat))
		}
		sort.Ints(f)
		sort.Ints(s)
		if equal(f, s) == true {
			fmt.Println("The strings are anagrams (equal length and equivalent letters)")
		} else {
			fmt.Println("The strings are not anagrams (equal length but unequivalent letters)")
		}
	}
}

func equal(one, two []int) bool {
	if len(one) != len(two) {
		return false
	}
	for i := range one {
		if one[i] != two[i] {
			return false
		}	
	}
	return true
}