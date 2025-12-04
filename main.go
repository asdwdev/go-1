// 3️⃣ Membuat slice dari array (opsional)
package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:4] // index 1 sampai 3
	fmt.Println(s)
}
