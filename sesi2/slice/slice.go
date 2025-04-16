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
	// copy func
	cp := copy(buah1, buah2)
	fmt.Printf("%#v \n", buah1)
	fmt.Printf("%#v \n", cp)

	//slicing
	var lookism = []string{"joggun", "junggo", "yohan", "hyungseok"}

	var lookism1 = lookism[1:4]
	fmt.Printf("%#v\n", lookism1)

	var lookism2 = lookism[3:]
	fmt.Printf("%#v\n", lookism2)

	var lookism3 = lookism[:3]
	fmt.Printf("%#v\n", lookism3)

	var lookism4 = lookism[:]
	fmt.Printf("%#v\n", lookism4)

	//slicing append
	lookism = append(lookism[:3], "Zin")
	fmt.Printf("%#v\n", lookism)

	//backing
	var lookismBacking = lookism[2:4]
	lookismBacking[0] = "God Dog"
	fmt.Printf("%#v\n", lookismBacking)

	//cap func
	//cap untuk mengetahui isi dari array maupun slice
	//len untuk mengetahui panjang dari array maupun slice
	var valorant = []string{"chyper", "waylay", "killjoy", "reyna"}
	fmt.Println("Valorant cap:", cap(valorant)) //4
	fmt.Println("Valorant len:", len(valorant)) //4
	fmt.Println(strings.Repeat("#", 30))

	var valorant1 = valorant[1:]
	fmt.Println("Valorant cap:", cap(valorant1)) //3
	fmt.Println("Valorant len:", len(valorant1)) //3
	fmt.Println(strings.Repeat("#", 30))

	var valorant2 = valorant[0:3]
	fmt.Println("Valorant cap:", cap(valorant2)) //4
	fmt.Println("Valorant len:", len(valorant2)) //3
	fmt.Println(strings.Repeat("#", 30))

	//newbacking
	roko := []string{"malboro", "garpit", "surya", "mild"}
	newRoko := []string{}

	newRoko = append(newRoko, roko[0:2]...)

	roko[0] = "esse"
	fmt.Println("roko: ", roko)
	fmt.Println("new roko: ", newRoko)
}
