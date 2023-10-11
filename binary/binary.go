package main

func binary(a []int, x int) (int, int) {
	size := len(a)
	start := 0
	middle := (start + size) / 2
	for start < size {
		if a[middle] == x {
			return a[middle], middle
		} else if a[middle] > x {
			size = middle - 1
		} else if a[middle] < x {
			size = middle + 1
		}
	}
	return -1, -1
}
