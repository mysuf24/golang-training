package main

import "fmt"

func main() {
	var number [4]int
	number = [4]int{1, 2, 3, 4}

	// var string = [3]string{"yusuf", "puccup", "daniel"}

	//verb %#v untuk memformat tipe data array agar kita juga dapat melihat panjang dari arraynya
	fmt.Printf("%#v\n", number)
	// fmt.Printf("%#v\n", string)

	// modifikasi index
	var buah = [3]string{"anggur", "semangka", "jeruk"}
	buah[0] = "grape"
	buah[1] = "watermelon"
	buah[2] = "orange"

	fmt.Printf("%#v\n", buah)

}
