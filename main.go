// latihan bagian 5
package main

import "fmt"

func kelilingPersegi(sisi int) int {
	return 4 * sisi
}

func luasPersegi(sisi int) int {
	return sisi * sisi
}

func hitung(sisi int) (int, int) {
	return kelilingPersegi(sisi), luasPersegi(sisi)
}

func main() {
	var sisi int
	fmt.Print("masukkan sisi: ")
	fmt.Scanln(&sisi)

	keliling, luas := hitung(sisi)
	fmt.Println("keliling:", keliling)
	fmt.Println("luas:", luas)
}
