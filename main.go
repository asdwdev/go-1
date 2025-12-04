// 4️⃣ Multiple Return Values (fitur khas Go)
// sangat sering dipakai.
package main

import "fmt"

func bagi(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa bagi 0")
	}
	return a / b, nil
}

func main() {
	hasil, err := bagi(10, 0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("hasil:", hasil)
}
