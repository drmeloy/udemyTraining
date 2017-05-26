package main

import "fmt"

func madz(n int) (int, bool){
	return n/2, n%2 == 0
}

func main(){
	h, even := madz(5)
	fmt.Println(h, even)
}