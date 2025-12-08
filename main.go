// 3️⃣ Membuat error custom (lebih rapi)
package main

import "errors"

var ErrSalahInput = errors.New("input tidak valid")

func main() {
	// contoh pemaiakan
	println(ErrSalahInput.Error())
}
