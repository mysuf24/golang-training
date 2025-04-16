package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// byte by byte
	city := "Tangerang"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte : %d\n", i, city[i])
	}

	var city1 []byte = []byte{84, 97, 110, 103, 101, 114, 97, 110, 103}
	fmt.Println(string(city1))

	//rune by rune
	str1 := "makan"
	str2 := "mânca"

	fmt.Printf("str1 byte lenght : %d\n", len(str1))
	fmt.Printf("str2 byte lenght : %d\n", len(str2))
	//rune count in string
	fmt.Printf("str1 char lenght : %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 char lenght : %d\n", utf8.RuneCountInString(str2))

	for i, s := range str2 {
		fmt.Printf("index: %d, string = %s\n", i, string(s))
	}
}
