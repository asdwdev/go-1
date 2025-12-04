// 1️⃣ ARRAY (fixed size)

// Array punya ukuran tetap. Jarang digunakan secara langsung.
package main

func main() {
	var a [3]int = [3]int{1, 2, 3}
	// atau
	// b := [5]string{"a", "b", "c", "d", "e"}

	println(a[0])

	// 	Ukuran array bagian dari tipenya:
	// [3]int ≠ [4]int
}
