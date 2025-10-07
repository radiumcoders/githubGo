package arrayslices

import "fmt"

func ArraySlices() {
	// Array
  arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Slice
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", slice)

	fmt.Printf("Array Type: %T Array Capacity: %d\n", arr, cap(arr))
	fmt.Printf("Slice Type: %T Slice Capacity: %d\n", slice, cap(slice))

	// slice with custom capacity 
	customSlice := make([]int, 3, 5)
	fmt.Printf("Custom Slice: %v, Length: %d, Capacity: %d\n", customSlice, len(customSlice), cap(customSlice))
}