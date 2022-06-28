package main

import "fmt"

var outs = "outs outside"
var pout = &outs

func main() {
	//mat := 34
	//var wale = "Second"
	outs := "outs inside"
	wale := "wale inside"
	var pin, pwale = &outs, &wale
	fmt.Println(pin , pwale , pin == pout)
}