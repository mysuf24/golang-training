package main

import (
	"fmt"
	"strings"
)

func main() {
	var ego = []string{"yusuf", "daniel", "ucup"}
	_ = ego

	fmt.Printf("%#v \n", ego)
	fmt.Println(strings.Repeat("#", 25))

	//make func
	var ego1 = make([]string, 3)
	_ = ego1

	ego1[0] = "yusuf"
	ego1[1] = "daniel"
	ego1[2] = "usup"

	fmt.Printf("%#v \n", ego1)
	fmt.Println(strings.Repeat("#", 25))

	// append func
	var ego2 = make([]string, 3)

	ego2 = append(ego2, "yusuf", "daniel", "usup")

	fmt.Printf("%#v \n", ego2)
	fmt.Println(strings.Repeat("#", 25))

	// append with ellipsis
	var buah1 = []string{"apel", "mangga", "jeruk"}
	var buah2 = []string{"anggur", "salak", "nanas"}

	buah1 = append(buah1, buah2...)

	fmt.Printf("%#v", buah1)

}
