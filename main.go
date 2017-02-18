package main

import "fmt"

func fib(n int) int {
	var ret int
	if n <= 0 {
		// error val
		ret = -1
	} else if n == 1 {
		ret = 0
	} else if n == 2 || n == 3 {
		ret = 1
	} else {
		ret = fib(n-1) + fib(n-2)
	}
	return ret
}

func fib_(n int) int {
	var p, q, ret int
	p, q = 0, 1
	for i := 3; i <= n; i++ {
		p, q = q, p+q
	}

	if n <= 0 {
		// error val
		ret = -1
	} else if n == 1 {
		ret = 0
	} else if n == 2 || n == 3 {
		ret = 1
	} else {
		ret = q
	}
	return ret
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(fib_(i))
	}
}
