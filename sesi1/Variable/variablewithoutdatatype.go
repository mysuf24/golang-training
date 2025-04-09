package main

import "fmt"

func main() {
	/* var name = "Yusuf"
	var age = 23 */

	//short declaration yang berarti kita harus langsung memberikan nilai kenapa variable tersebut jika tidak akan terjadi error
	name := "Yusuf"
	age := 23
	
	//print F bergantung pada flag yang akankita pakai, contoh verb% T kegunaan nya untuk menentukan tipe data yang akan digunakan
	fmt.Printf("%T, %T", name, age)
}