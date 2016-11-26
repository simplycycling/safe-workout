package main

import "fmt"

func main() {
	var temp int
	_, err := fmt.Scanf("%d", &temp)
	var	dewp int
	_, err := fmt.Scanf("%d", &dewp)
	
	if temp + dewp >= 150 {
		fmt.Println("Tejas Richard says do it indoors.")
	} else {
		fmt.Println("Tejas Richard says go do it outdoors!")
	}
}