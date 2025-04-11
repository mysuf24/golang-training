package main

import "fmt"

func main() {
	// 2 lebih kecil 3? benar 
	var firstCondition bool = 2 < 3
	// 2 lebih besar dari 3? salah
	var secondCondition bool = 2 > 3
	// yusuf sama dengan yusuf? benar
	var thirdCondition bool = "yusuf" == "yusuf"
	// yusuf sama dengan mochammad? salah
	var fourthCondition bool = "yusuf" == "mochammad"
	// 10 tidak sama dengan 2.4? benar
	var fifthCondition bool = 10 != 2.4
	// 24 kurang dari atau sama dengan 23? salah
	var sixthCondition bool = 24 <= 23

	fmt.Println(" first condition:", firstCondition)
	fmt.Println(" second condition:", secondCondition)
	fmt.Println(" third condition:", thirdCondition)
	fmt.Println(" fourth condition:", fourthCondition)
	fmt.Println(" fifth condition:", fifthCondition)
	fmt.Println(" sixth condition:", sixthCondition)
}
