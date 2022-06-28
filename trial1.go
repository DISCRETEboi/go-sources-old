//This program takes in text from user and rotates it right or left as specified
package main

import "fmt"

func main() {
	var args string
	var rot byte
	var dir string
	fmt.Print("Enter the text to rotate (without spaces) --> ")
	fmt.Scanf("%s" , &args)
	fmt.Print("Enter the amount to rotate by (positive integer) --> ")
	fmt.Scanf("%d" , &rot)
	fmt.Print("Enter the direction to rotate ('right' or 'left') --> ")
	fmt.Scanf("%s" , &dir)
	l := len(args)
	m1 := []byte(args)
	if dir == "right" {
		val := byte(l) - rot
		m2 := make([]byte, rot, l)
		copy(m2, m1[val:])
		m2 = append(m2, m1[:val]...)
		var m3 string = string(m2)
		fmt.Println("The rotated string is" , m3)
	} else if dir == "left" {
		m2, m3 := make([]byte, rot), make([]byte, byte(l) - rot)
		copy(m2, m1[:rot])
		copy(m3, m1[rot:])
		m4 := string(append(m3, m2...))
		fmt.Println("The rotated string is" , m4)
	} else {
		fmt.Println("Make sure your arguments are correct!")
	}
}