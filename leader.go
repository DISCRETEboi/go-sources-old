//Qamardeen

package main

import (
	"fmt"
	"sort"
	"strings"
	"os"
)

var a, b, c, d, e string
var list0 []string
var in = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	fmt.Print("Enter the name of the first participant in *UPPER CASE*:--> ")
	fmt.Scan(&a)
	fmt.Print("Enter the name of the second participant in *UPPER CASE*:--> ")
	fmt.Scan(&b)
	fmt.Print("Enter the name of the third participant in *UPPER CASE*:--> ")
	fmt.Scan(&c)
	fmt.Print("Enter the name of the fourth participant in *UPPER CASE*:--> ")
	fmt.Scan(&d)
	fmt.Print("Enter the name of the fifth participant in *UPPER CASE*:--> ")
	fmt.Scan(&e)
	list0 = []string{a, b, c, d, e}
	var test bool
	for _, mat := range list0 {
		for _, _ = range mat {
			test = strings.ContainsAny(string(mat), in)
			if test == false {
				fmt.Println("The input contains contrabands and the program exits!")
				os.Exit(0)
			}
		}
	}	
}

func main() {
	list := make([]string, 0)
	for _, mat := range list0 {
		astr := ""
		for _, man := range mat {
			for _, mad := range in {
				if man == mad {
					astr = astr + fmt.Sprintf("%c" , man)
				}
			}
		}
		list = append(list, astr)
	}
	m1 := make([]int, 0)
	for _, str := range list {
		m1 = append(m1, len(str))
	}
	orig := make([]int, len(list))
	copy(orig, m1)
	sort.Ints(m1)
	maxCount := m1[len(m1)-1]
	var m2 []string
	for _, mat := range list {
		if len(mat) == maxCount {
			m2 = append(m2, mat)
		}
	}
	sort.Strings(m2)
	ans := m2[0]
	fmt.Println("The president is" , ans)
}