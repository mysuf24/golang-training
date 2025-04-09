package main

import "fmt"

func main() {
	//tipe data string adalah nilainya di apit oleh tanda petik dua (“”)
	var pesan string = "Hallo"
	//Keistimewaan string yang dideklarasikan menggunakan backtics adalah membuat semua karakter didalamnya tidak di
escape, termasuk \n, tanda petik dua dan tanda petik satu, baris baru, dan lainnya. Semua akan terdeteksi sebagai string.
	var pesan1 = `ini menggunakan backtics`
	fmt.Print(pesan)
	fmt.Print(pesan1)
}