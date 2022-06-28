package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	str := ""
	sep := ""
	for _, val := range args {
		str += sep + val
		sep = " "
	}
	fmt.Println(str)
}