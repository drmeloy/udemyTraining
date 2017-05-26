package main

import "fmt"

func main(){
	var sum int
	for a := 0; a < 1000; a++ {
		if a % 3 == 0 {
			sum += a
		}else if a % 5 == 0{
			sum += a
		}
	}
	fmt.Println(sum)
}