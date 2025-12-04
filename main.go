// If, For, Switch (Kontrol Alur)
package main

import "fmt"

func main() {
	// 	2️⃣ FOR — Satu-satunya loop di Go

	// Go cuma punya 1 jenis loop yaitu for.
	// Tapi bisa dipakai dalam banyak gaya.

	// 2.3 For tanpa kondisi (infinite loop)
	for {
		fmt.Println("loop terus")
	}
	// Biasanya dipakai untuk server atau event loop.
}
