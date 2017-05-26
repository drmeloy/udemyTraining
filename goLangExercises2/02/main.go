package main

import "fmt"

func main(){
	madz := func(n int)(int, bool){
		return n/2, n%2 == 0
	}
	fmt.Println(madz(5))
}
