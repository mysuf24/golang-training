package main

import "fmt"

func main() {
	var value = 2 + 2*3
	var value1 = (2 + 2) * 3
	// pada bahasa Go syntax kita akan dibaca dari kiri ke kanan tetapi karena ada simbol * atau perkalian maka go akan melakukan perkalian terlebih dahulu jika ingin melakukan pertambahan terlebih dahulu tinggal di tambahkan simbol ()
	fmt.Println("ini akan melakukan perkalian terlebih dahulu: ", value)
	fmt.Println("ini akan melakukan pertambahan terlebih dahulu: ", value1)
}
