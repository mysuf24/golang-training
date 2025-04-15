package main

import "fmt"

func main() {
	// cara pertama
	// jika i sama dengan 0, i lebih kecil dari 4, i = i + 1 jika tidak di tambahkan akan menimbulakan infinite loop. maka hasil nya akan terjadi pengulangan sebanyak 4x
	for i := 0; i < 4; i++ {
		fmt.Println("angka", i)
	}

	// cara ke 2
	var i = 0

	for i < 3 {
		fmt.Println("Angka", i)
		i++
	}
}
