package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	str1, str2 := swap("Dhabi", "Abu")

	fmt.Println(str1, str2)
}
