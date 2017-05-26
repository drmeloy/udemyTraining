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

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

ANSWER: 233168

P.S. While this was question one on Project Euler, I will try another problem to send to Todd because we had already done this one, and if I chose only this one to send it would be a cop out.
 */
