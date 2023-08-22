package main

import "fmt"

func main() {
	var i int
	var n int
	var x int
	var arr [100]int

	fmt.Println("Enter the size")
	fmt.Scan(&n)

	fmt.Println("Enter the array")
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Enter the key")
	fmt.Scan(&x)

	for i = 0; i < n; i++ {
		if arr[i] == x {
			fmt.Println("True")
			break
		}
	}
	if i == n {
		fmt.Println("False")
	}
}
