package main

import "fmt"

func main() {
	var score int = 90
	// 无需使用 break
	switch score {
	case 90:
		fmt.Println("90")
	case 80:
		fmt.Println("80")
	case 70:
		fmt.Println("70")
	case 60:
		fmt.Println("60")
	case 50:
		fmt.Println("50")
	}

	var s string
	fmt.Println(s)

	arrays := []string{"a", "b", "c", "d"}
	for index, value := range arrays {
		fmt.Printf("%d ------> %s", index, value)
		fmt.Println()
	}
}
