package main

import "fmt"
import "res1"
import "os"

func main() {
	var str string
	fmt.Print("Enter the word >> ")
	fmt.Scanf("%s" , &str)
	legal := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if res1.Subset(res1.Str2ints(legal), res1.Str2ints(str)) == false {
		fmt.Println("Error: Make sure your arguments are letters in capital!")
		os.Exit(0)
	}
	var s1 []int
	for _, mat := range str {
		s1 = append(s1, int(mat))
	}
	word := []int{s1[0]}
	for _, mat := range s1[1:] {
		if mat >= word[0] {
			word = res1.Insert(word, int(mat), 0)
		} else {
			word = append(word, int(mat))
		}
	}
	fmt.Printf("%c" , word)
}