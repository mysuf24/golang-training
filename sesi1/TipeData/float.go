package main

import "fmt"

func main() {
	var decimalNumber float32 = 3.63
	
	// verb %f pada fungsi fmt.Printf untuk memformat nilai float tersebut
	fmt.Printf("decimal Number: %f\n", decimalNumber)
	// Jika kita ingin mengecilkan digit desimalnya, maka kita dapat menggunakan format verb seperti ( %.nf )
	fmt.Printf("decimal Number: %.3f\n", decimalNumber)
}