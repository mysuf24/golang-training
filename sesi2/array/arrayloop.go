package main

import (
	"fmt"
	"strings"
)

func main() {
	var ego = [3]string{"yusuf", "daniel", "usup"}
	//cara pertama
	for i, v := range ego {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	//cara kedua

	// fungsi len untu mendapatkan panjang dari array nya dengan cara penulisan len(<nama array>).
	for i := 0; i < len(ego); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, ego[i])
	}
}
