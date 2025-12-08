// custom error + wrapping + errors.Is
package main

import (
	"errors"
	"fmt"
)

type SalahInputError struct {
	Pesan string
}

func (e SalahInputError) Error() string {
	return e.Pesan
}

// 2. Fungsi yang mengembalikan error terbungkus (fmt.Errorf("%w"))
func cekUmur(umur int) error {
	if umur < 18 {
		err := SalahInputError{Pesan: "umur harus >= 18"}
		return fmt.Errorf("validasi gagal: %w", err) // wrap error
	}
	return nil
}

// 3. Cara mengecek apakah error tersebut adalah SalahInputError
// Gunakan errors.As (lebih cocok daripada Is untuk tipe struct).
func main() {
	err := cekUmur(15)
	if err != nil {
		var inputErr SalahInputError

		// cek apakah err mengandung SalahInputError (meskipun di wrap)
		if errors.As(err, &inputErr) {
			fmt.Println("terjadi error input:", inputErr.Pesan)
			return
		}

		// fallback error biasa
		fmt.Println("error lain:", err)
	}
}

// errors.Is() cocok untuk sentinel error (var errSomething = errors.New("…"))
// Tapi untuk custom struct error, yang benar pakai errors.As().
