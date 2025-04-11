package main

import "fmt"

func main() {
	var benar = true
	var salah = false

	// operator dan dan : jika kedua nya sama maka jawaban nya true
	var salahDanBenar = benar && benar
	fmt.Printf("salah dandan benar \t(%t) \n", salahDanBenar)

	// operator or : jika salah satu nya berbeda maka jawaban nya salah
	var salahAtauBenar = salah || salah
	fmt.Printf("salah atau benar \t(%t) \n", salahAtauBenar)

	// keterbalikan atau lawan dari
	var salahTerbalik = !salah
	fmt.Printf("salah terbalik \t(%t) \n", salahTerbalik)
}
