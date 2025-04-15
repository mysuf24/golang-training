package main

import "fmt"

func main() {
outerLoop:
	// i dimulai dari 0, i lebih kecil dari 3, i= i + 1
	for i := 0; i < 3; i++ {
		// ini akan melakukan print sebanyak 3x
		fmt.Println("Perulangan ke = ", i+1)
		// j di mulai dari 0, lalu j lebih kecil dari 3, j= j+1
		// jika i sama dengan 2 maka berhenti, lalu melakukan perulangan ulang ke luar
		// yg artinya j akan berhenti sampai perulangan ke 2
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Println(j, " ")
		}
		fmt.Println("\n")
	}
}
