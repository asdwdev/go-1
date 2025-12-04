// named return values
// Go bisa memberi nama pada return value:
package main

import "fmt"

func persegi(s int) (luas int, keliling int) {
	luas = s * s
	keliling = 4 * s
	return // return tanpa nilai -> pakai nilai bernama
}

func main() {
	luas, keliling := persegi(5)
	fmt.Println("luas persegi:", luas)
	fmt.Println("keliling persegi:", keliling)
}
