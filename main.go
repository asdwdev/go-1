// 4️⃣ MAP (key-value)

package main

import "fmt"

func main() {
	// Contoh map string → int:
	m := map[string]int{
		"apel":   3,
		"pisang": 5,
	}
	fmt.Println(m)

	// menambahkan key
	m["jeruk"] = 7
	fmt.Println(m)

	// looping map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Cek apakah key ada:
	v, ok := m["apels"]
	if ok {
		fmt.Println("ada:", v)
	} else {
		fmt.Println("key tidak ditemukan")
	}

}
