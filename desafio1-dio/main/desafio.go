package main

import "fmt"

func main() {

	var arr []int

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			arr = append(arr, i)
		}
	}

	fmt.Println(arr)
}
