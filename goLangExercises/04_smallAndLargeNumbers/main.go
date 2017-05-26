package main

import "fmt"

func main(){
	var small int
	var large int
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&small)
	fmt.Print("Thank you. Now please enter a large number: ")
	fmt.Scan(&large)
	fmt.Println("Thank you.", large, "divided by", small, "gives you a remainder of", small%large)
}
