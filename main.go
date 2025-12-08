// ✅ 2. Custom error tapi minimalis (pakai struct, tapi ringkas)
package main

type SalahInputError string

func (e SalahInputError) Error() string {
	return string(e)
}

func main() {
	err := SalahInputError("input salah banget")
	println(err.Error())
}
