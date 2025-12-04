// named return values
// Go bisa memberi nama pada return value:
package main

import "fmt"

func luasPersegi(s int) (luas int) {
	luas = s * s
	return // return tanpa nilai -> pakai nilai bernama
}

func main() {
	hasil := luasPersegi(2)
	fmt.Println(hasil)
}
