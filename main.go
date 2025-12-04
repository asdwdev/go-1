// 4️⃣ MAP (key-value)

package main

import "fmt"

func main() {
	// Contoh map string → int:
	m := map[string]int{
		"apel":   3,
		"pisang": 5,
	}
	fmt.Println(m)

	// delete element map
	delete(m, "apel")
	fmt.Println(m)

}
