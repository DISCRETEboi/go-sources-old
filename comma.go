//Written by qamardeen abdul-mateen aka discrete_boi
package main

import "fmt"

func main() {
	number := "232750689569623565049355717"
	var n int
	long := len(number)
	rem := long % 3
	quo := int(long / 3)
	if long <= 3 {
		fmt.Println(number)
		return 
	} else {
		n = quo
	}
	m := long - 3
	ini := ""
	for i := 1; i <= n; i++ {
		if i == n && rem == 0 {
			ini = number[m:m + 3] + ini
			m -= 3
		} else {
			ini = "," + number[m:m + 3] + ini
			m -= 3
		}
	}
	ini = number[:rem] + ini
	fmt.Println(ini)
}