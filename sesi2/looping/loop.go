package main

import "fmt"

func main() {
	// jika i sama dengan 0, i lebih kecil dari 4, i = i + 1. maka hasil nya akan terjadi pengulangan sebanyak 3x
	for i := 0; i < 4; i++ {
		fmt.Println("angka", i)
	}
}
