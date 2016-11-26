package main

import "fmt"

func main() {
	var temp int
	_, err := fmt.Scanf("What is the current temperature? %d", &temp)
	var	dewp int
	_, err := fmt.Scanf("And what is the current dew point? %d", &dewp)
	
	if temp + dewp >= 150 {
		fmt.Println("Tejas Richard says do it indoors.")
	} else {
		fmt.Println("Tejas Richard says go do it outdoors!")
	}
}