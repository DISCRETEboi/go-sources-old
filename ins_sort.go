package main

import "fmt"
import r1 "res1"

func main() {
	arg := [...]int{2,6,4,9,4,1,3,1,5,7,4,7,5,9,2,6}
	ans := make([]int, len(arg))
	if arg[1] >= arg[0] {
		ans[0], ans[1] = arg[0], arg[1]
	} else {
		ans[0], ans[1] = arg[1], arg[0]
	}
	for _, mat := range arg[2:] {
		if mat < arg[0] {
			ans = r1.Insert(ans, mat, 0)
			fmt.Println(ans)
			break
		}
		for x, y := 0, 1; true; x, y = x+1, y+1 {
			if mat >= arg[x] && mat <= arg[y] {
				ans = r1.Insert(ans, mat, y)
				break
			}
		}
		fmt.Println(ans)
		fmt.Println(len(arg) , len(ans))
	}
	ans = r1.RemoveAll(ans, 0)
	fmt.Println(ans)
}
