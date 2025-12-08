// Custom error dengan data tambahan (tetap simpel)
package main

import "fmt"

type SalahInputError struct {
	Field string
}

func (e *SalahInputError) Error() string {
	return "input salah pada field:" + e.Field
}

func main() {
	err := &SalahInputError{Field: "username"}
	fmt.Println(err.Error())
}
