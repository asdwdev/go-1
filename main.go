// latihan
package main

import (
	"errors"
	"fmt"
	"math"
)

func akar(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("tidak bisa akar bilangan negatif")
	}

	return math.Sqrt(x), nil
}

func main() {
	var angka float64
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scanln(&angka)

	v, err := akar(angka)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("hasil:", v)
}
