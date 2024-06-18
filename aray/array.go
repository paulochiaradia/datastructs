package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 9, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println("Printing elements ", arr[i])
	}

	for i, value := range arr {
		fmt.Println("index", i, " range", value)
	}
}
