package main

import "fmt"

func countdown(num int) {
	for i := num; i > 0; i-- {
		fmt.Println(i)
	}
}

// function with return value
func sum(a, b int) int {
	return a + b
}

func main(){
	countdown(10)
	sum(2,3)
}