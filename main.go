// latihan bagian 6
package main

import "fmt"

func main() {
	// === angka ===
	angka := []int{1, 2, 3, 4, 5}

	angka = append(angka, 6, 7)

	for _, v := range angka {
		fmt.Println(v)
	}

	// === buah ===
	m := map[string]int{
		"apel":   3,
		"mangga": 5,
		"jeruk":  2,
	}

	m["pisang"] = 10

	for k, v := range m {
		fmt.Println(k, v)
	}

}
