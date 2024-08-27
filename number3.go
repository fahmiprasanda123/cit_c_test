package main

import (
	"fmt"
)

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func Mati(b *Bebek) {
	b.hidup = false
}

func Terbang(b *Bebek) {
	if b.energi > 0 && b.hidup && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			Mati(b)
		}
	}
}

func Makan(b *Bebek) {
	if b.hidup {
		b.energi += 1
	}
}

func main() {
	bebek := Bebek{
		energi:       3,
		hidup:        true,
		bisaTerbang:  true,
		suaraTerbang: "Kwek!",
	}

	// Bebek terbang
	Terbang(&bebek)
	Terbang(&bebek)
	Terbang(&bebek)

	fmt.Println("Energi bebek setelah terbang:", bebek.energi)
	fmt.Println("Bebek masih hidup:", bebek.hidup)

	// Bebek makan
	Makan(&bebek)
	fmt.Println("Energi bebek setelah makan:", bebek.energi)
}
