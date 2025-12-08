package main

import "fmt"

// 1️⃣ Apa itu interface di Go?

// Interface = kumpulan method yang mendefinisikan perilaku.
type Speaker interface {
	Speak() string
}

// 🎯 “Tipe apa pun yang punya method Speak() string otomatis dianggap sebagai Speaker.”
// implementasi interface (otomatis!)

type Dog struct{}

func (d Dog) Speak() string {
	return "guk!"
}

// pemakaian interface
func SaySomething(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	d := Dog{}
	SaySomething(d)
}
