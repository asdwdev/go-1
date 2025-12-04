// If, For, Switch (Kontrol Alur)
package main

import "fmt"

func main() {
	// 3️⃣ SWITCH — Lebih sederhana dari if bertingkat
	// Go punya switch yang bersih dan sederhana:

	hari := "mon"

	switch hari {
	case "mon":
		fmt.Println("senin")
	case "tue":
		fmt.Println("selasa")
	default:
		fmt.Println("hari lain")
	}

	// Di Go, switch otomatis break — tidak perlu break.

}
