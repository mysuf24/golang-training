package main

import "fmt"

func main() {
	var score = 1

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("coba lagi yaa")
		} else {
			fmt.Println("ayo kamu bisa")
			if score == 0 {
				fmt.Println("kamu kurang usaha")
			}
		}
	}
}
