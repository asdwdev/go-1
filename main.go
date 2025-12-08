// 1️⃣ Error basic: return error

// Go punya tipe built-in error.
package main

import (
	"errors"
	"fmt"
)

func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak bisa bagi dengan nol")
	}
	return a / b, nil
}

func main() {
	hasil, err := bagi(10, 0)
	if err != nil {
		fmt.Println("terjadi error:", err)
		return
	}
	fmt.Println("hasil:", hasil)
}

// result, err := something()
// if err != nil {
//     // tangani error
// }
