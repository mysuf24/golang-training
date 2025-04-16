package main

import (
	"fmt"
	"strings"
)

func main() {
	var person map[string]string // deklarasi
	person = map[string]string{} // inisialisasi

	person["nama"] = "Yusuf"
	person["umur"] = "23"
	person["alamat"] = "Jalan Pidana"

	fmt.Println("nama :", person["nama"])
	fmt.Println("umur :", person["umur"])
	fmt.Println("alamat :", person["alamat"])
	fmt.Println(strings.Repeat("#", 25))

	//cara ke 2
	var person1 map[string]string
	person1 = map[string]string{
		"nama":   "Usup",
		"umur":   "12",
		"alamat": "bambu apus",
	}

	fmt.Println("nama :", person1["nama"])
	fmt.Println("umur :", person1["umur"])
	fmt.Println("alamat :", person1["alamat"])
	fmt.Println(strings.Repeat("#", 25))

	//looping map
	for key, value := range person {
		fmt.Println(key, ":", value)
	}
	fmt.Println(strings.Repeat("#", 25))

	//deleting value
	fmt.Println("Sebelum dihapus:", person1)
	delete(person1, "alamat")
	fmt.Println("Sesudah dihapus:", person1)

	//deleting a value
	value, exist := person["umur"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesnt exist")
	}

	delete(person, "umur")

	value, exist = person["umur"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	//map combine with slice
	var ego = []map[string]string{
		{"nama": "yusuf", "umur": "23"},
		{"nama": "usup", "umur": "12"},
		{"nama": "daniel", "umur": "18"},
	}

	for i, alter := range ego {
		fmt.Printf("index: %d, nama: %s, umur: %s\n", i, alter["nama"], alter["umur"])
	}
}
