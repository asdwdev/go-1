// fugnsi dengan parameter
package main

import "fmt"

func sapa(nama string) {
	fmt.Println("halo,", nama)
}

// parameter bisa lebih dari satu
func tambah(a int, b int) {
	fmt.Println(a + b)
}

// atau lebih ringkas:
func kurang(a, b int) {
	fmt.Println(a - b)
}

func main() {
	sapa("nnl")
	tambah(2, 3)
	kurang(3, 2)
}
