package main

import "fmt"

func main() {
	var score = 9

	//tanpa operator
	switch score {
	case 8:
		fmt.Println("sempurna")
	case 7:
		fmt.Println("bagus")
	default:
		fmt.Println("cukup \n")
	}

	//dengan operator
	switch {
	case score == 8:
		fmt.Println("sempurna")
	case (score < 8) && (score > 3):
		fmt.Println("cukup")
	default: //jika melewati batas score maka akan dihitung default
		{
			fmt.Println("belajar lagi")
			fmt.Println("kamu kurang giat")
		}
	}

	//fallthrough
	switch {
	case score == 8:
		fmt.Println("sempurna")
	case (score < 8) && (score > 3):
		fmt.Println("cukup")
		fallthrough
	case score < 5:
		fmt.Println("ayo belajar lagi")
	default: //jika melewati batas score maka akan dihitung default
		{
			fmt.Println("tolong belajar lagi")
			fmt.Println("kamu kurang giat")
		}
	}
}
