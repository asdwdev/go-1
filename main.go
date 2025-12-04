// If, For, Switch (Kontrol Alur)
package main

import "fmt"

func main() {
	// 	1.1 IF dengan kondisi di dalamnya (idiom Go)

	// Go memperbolehkan membuat variabel langsung di dalam if.
	if x := 10; x > 5 {
		fmt.Println("lebih besar")
	}
}
