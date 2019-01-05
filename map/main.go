package main

import "fmt"

func main() {
	x := map[string]string{"a": "aa", "b": "bb", "c": "cc"}

	for k, v := range x {
		fmt.Print(k + " ---> " + v + "\n")
	}
}
