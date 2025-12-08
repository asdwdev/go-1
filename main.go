// 1. Custom Error Struct
package main

type SalahInputError struct {
	Pesan string
}

func (e SalahInputError) Error() string {
	return e.Pesan
}

// ✔️ 2. Cara Membuat dan Mengembalikan Custom Error

// Biasanya custom error dipakai di fungsi yang butuh validasi:
func cekUmur(umur int) error {
	if umur < 0 {
		return SalahInputError{Pesan: "umur tidak boleh negatif"}
	}
	if umur < 18 {
		return SalahInputError{Pesan: "umur harus minimal 18 tahun"}
	}
	return nil
}

// pemakaian di main()
func main() {
	err := cekUmur(15)
	if err != nil {
		// cek apakah errornya tipe custom kita
		switch e := err.(type) {
		case SalahInputError:
			println("error input:", e.Pesan)
		default:
			println("error lain:", err.Error())
		}
	}
}
