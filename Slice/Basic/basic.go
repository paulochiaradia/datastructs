package main

import "fmt"

func main() {
	slice := []int{1, 3, 5, 6}
	fmt.Println("Capacity ", cap(slice))
	slice = append(slice, 8)
	fmt.Println("Capacity ", cap(slice))
	fmt.Println("Lenght", len(slice))
}
