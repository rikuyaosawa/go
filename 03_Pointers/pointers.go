package main

import "fmt"

func main() {
	hisAge := 24
	herAge := 22

	pointerAge := &herAge

	fmt.Println("He have", getRemainingYearsTill80(hisAge), "years till he is 80.")
	fmt.Println("She have", getRemainingYearsTill100(*pointerAge), "years till she is 1000.")
	setAgeTo50(&hisAge)
	fmt.Println("Oh! His age is now", hisAge)
}

func getRemainingYearsTill80(age int) int {
	return 80 - age
}

func getRemainingYearsTill100(age int) int {
	return 100 - age
}

func setAgeTo50(age *int) {
	*age = 50
}
