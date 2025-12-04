// If, For, Switch (Kontrol Alur)
package main

import "fmt"

func main() {
	// 4️⃣ range — Loop untuk slice, array, map, string

	// range sering digunakan untuk mengulang isi struktur data.

	// Untuk slice:
	angka := []int{10, 20, 30}

	for i, v := range angka {
		fmt.Println("index:", i, "nilai:", v)
	}

	// kalau hanya butuh nilai for _, v := range angka {
}
