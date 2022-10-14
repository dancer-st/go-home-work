package main

import "fmt"

func sliceCapacity(slice []int) ([]int) {
	var capacity []int
	var new_slice []int
	var last_capacity int
	last_capacity = cap(new_slice)

	//capacity = append(capacity, last_capacity)
	for i:=0; i<len(slice); i++ {
		new_slice = append(new_slice, slice[i])
		if (cap(new_slice) != last_capacity) {
			last_capacity = cap(new_slice)
			len := len(new_slice);
			capacity = append(capacity, len)
		}
		//fmt.Println(capacity)
	}

	return capacity
}

func main() {
	var slice []int
	for i:=0; i<100; i++ {
		slice = append(slice, i) // len=2, cap=2
	}

	fmt.Println(sliceCapacity(slice))


}
