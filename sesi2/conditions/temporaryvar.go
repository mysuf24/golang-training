package main

import "fmt"

func main() {
	var currentYear = 2025

	// jika age = var tahun sekarng yaitu 2025 - 2002
	// lalu terdapat kondisi jika age lebih kecil dari 21 maka hasil nya sudah legal
	if age := currentYear - 2002; age < 21 {
		fmt.Println("belum legal")
	} else {
		fmt.Println("sudah legal")
	}
}