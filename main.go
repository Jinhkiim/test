package main

import "fmt"

func fib(n uint64) uint64 {
	if n <= 0 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(10))
}
