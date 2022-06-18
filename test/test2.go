package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func re(bool2 bool) bool {
	if bool2 == true {
		return re(bool2)
	}
	return false
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}

	ins := []bool{true, true, true, true, true, false}
	for i2, in := range ins {
		fmt.Println(i2, re(in))
	}
}
