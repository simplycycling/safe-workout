package main

import "fmt"

func main() {
	temp := fmt.Fscan()
	dewp := // current dew point
	
	if temp + dewp >= 150 {
		fmt.Println("Tejas Richard says do it indoors.")
	} else {
		fmt.Println("Tejas Richard says go do it outdoors!")
	}
}
