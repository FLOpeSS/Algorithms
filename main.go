package main

import "fmt"

func insertion_sort(array []int) []int {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
	return array
}

func main() {
	array := []int{4, 5, 2, 1}
	fmt.Println("original array", array)
	insertion_sort(array)
	fmt.Println("sorted array", array)

}
