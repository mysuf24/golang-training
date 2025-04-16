package main

import "fmt"

func main() {

	var a uint8 = 10
	var b byte //byte alias dari tipe data uint8

	b = a // tidak error, karena byte dan uint8 memiliki tipe data yang sama
	_ = b

	type second = uint

	var hour second = 3600
	fmt.Printf("hour type %T\n", hour) 
}
