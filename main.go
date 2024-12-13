package main

import (
	"flight/flight"
	flightmenu "flight/flightMenu"
	"fmt"
)

func main() {
	fmt.Println("Flight Booking Application")
	choice := -1
	for {
		fmt.Println("Enter your choice:- ")
		flightmenu.FirstMenu()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			{
				AdminFunc()
			}
		case 2:
			{
				UserFunc()
			}
		case 0:
			{
				return
			}
		}
	}
}

func UserFunc() {
	choice := -1
	for {
		flightmenu.UserMenu()
		fmt.Println("Enter your choice:- ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			{
				flight.UserFlightInput()
			}
		case 0:
			{
				return
			}
		}
	}

}

func AdminFunc() {
	choice := -1
	for {
		flightmenu.AdminMenu()
		fmt.Println("Enter your choice:- ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			{
				flight.AddFlight()
			}
		case 2:
			{
				flight.DeleteFlight()
			}

		case 0:
			{
				return
			}
		}
	}

}
