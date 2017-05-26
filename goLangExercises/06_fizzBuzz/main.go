package main

import "fmt"

func main(){
	for a := 0; a <= 100; a++ {
		if a % 15 == 0 {
			fmt.Println("FIZZBUZZ")
		} else if a % 5 == 0 {
			fmt.Println("BUZZ")
		} else if a % 3 == 0 {
			fmt.Println("FIZZ")
		} else {
			fmt.Println(a)}
	}
}