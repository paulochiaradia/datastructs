package main

import "fmt"

func twiceValue(slice []int) {
	for index, value := range slice {
		slice[index] = 2 * value
	}
}

func main() {
	slice := []int{1, 3, 5, 7}
	twiceValue(slice)
	for i, _ := range slice {
		fmt.Println("New Slice value", slice[i])
	}
}
