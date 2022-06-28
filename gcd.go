//This is my Greatest Common Divisor Algorithm in Go...
//Qamardeen Abdul-Mateen ... Copyright
package main

import "fmt"

func main() {
	m := 12345678
	n := 1234
	up, low := m, n
	rem := m % n
	count := 0
	for rem != 0 {
		fmt.Println("For GCD division count" , count+1 , "...")
		fmt.Println("The division of" , m , "by" , n , "has remainder" , rem , "not 0")
		fmt.Println("Then" , n , "will be divided by" , rem)
		fmt.Println("+-----------------------------------------------------------+")
		m = n
		n = rem
		rem = m % n
		count += 1
	}
	fmt.Println("Finally, the division of" , m , "by" , n , "gives a remainder of 0")
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println("The gcd of" , up , "and" , low , "is" , n)
}