package main

import "fmt"

func main() {
	//	untuk i = 1, i lebih kecil atau sama dengan 10, i++ artinya akan terus bertambah.
	// jika i%2 sama dengan 1 yang berarti ganjil akan melanjutkan
	// jika i lebih besar dari 8 maka berhenti
	// maka hasil yang di keluarkan akan genap dari 2 sd 8
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}
