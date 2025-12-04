// 2️⃣ SLICE (dynamic array) — INI YANG PALING PENTING

// Slice = struktur data yang fleksibel dan paling sering digunakan.
package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	fmt.Println(nums)

	// Slice bisa bertambah panjang:
	nums = append(nums, 40)
	fmt.Println(nums)

	// akses elemen
	fmt.Println(nums[0])

	// looping slice
	for i, v := range nums {
		fmt.Println(i, v)
	}
}
