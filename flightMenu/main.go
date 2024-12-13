package flightMenu

import "fmt"

func FirstMenu() {
	fmt.Println("1. Admin")
	fmt.Println("2. User")
	fmt.Println("0. Exit")
}

func AdminMenu() {
	fmt.Println("1. Add Flight")
	fmt.Println("2. Delete Flight")
	fmt.Println("0. Exit")
}

func UserMenu() {
	fmt.Println("1. Search Flight")
	fmt.Println("0. Exit")
}
