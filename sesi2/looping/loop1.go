package main

import "fmt"

func main() {
	var i = 0

	//looping menggunakan break untuk memberhentikan
	for {
		fmt.Println("Angka", i)

		i++
		if i == 3 {
			break
		}
	}
}
