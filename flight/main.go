package flight

import (
	"fmt"
	"sort"
	"time"
)

type ClassType int

const (
	Economy ClassType = iota
	Premium
	Business
	First
)

type Flight struct {
	Id            int
	From          string
	To            string
	FlightName    string
	Price         string
	Class         ClassType
	DepartureTime time.Time
	ArrivalTime   time.Time
	Date          time.Time
}

var Flights = []Flight{
	// Flights on 2024-12-15, same source and destination (Delhi to Mumbai)
	{Id: 1, From: "Delhi", To: "Mumbai", FlightName: "Air India AI101", Price: "₹7000", Class: Economy, DepartureTime: mustParseTime("06:45"), ArrivalTime: mustParseTime("09:15"), Date: mustParseDate("2024-12-15")},
	{Id: 2, From: "Delhi", To: "Mumbai", FlightName: "IndiGo 6E123", Price: "₹8500", Class: Business, DepartureTime: mustParseTime("07:30"), ArrivalTime: mustParseTime("10:00"), Date: mustParseDate("2024-12-15")},
	{Id: 3, From: "Delhi", To: "Mumbai", FlightName: "SpiceJet SG987", Price: "₹5000", Class: Premium, DepartureTime: mustParseTime("08:15"), ArrivalTime: mustParseTime("10:45"), Date: mustParseDate("2024-12-15")},
	{Id: 4, From: "Delhi", To: "Mumbai", FlightName: "Vistara VT456", Price: "₹6000", Class: Business, DepartureTime: mustParseTime("10:00"), ArrivalTime: mustParseTime("12:30"), Date: mustParseDate("2024-12-15")},
	{Id: 5, From: "Delhi", To: "Mumbai", FlightName: "GoAir G8123", Price: "₹4500", Class: Economy, DepartureTime: mustParseTime("12:00"), ArrivalTime: mustParseTime("14:30"), Date: mustParseDate("2024-12-15")},
	{Id: 6, From: "Delhi", To: "Mumbai", FlightName: "AirAsia AA234", Price: "₹5500", Class: Premium, DepartureTime: mustParseTime("14:00"), ArrivalTime: mustParseTime("16:30"), Date: mustParseDate("2024-12-15")},
	{Id: 7, From: "Delhi", To: "Mumbai", FlightName: "IndiGo 6E789", Price: "₹6500", Class: Business, DepartureTime: mustParseTime("15:00"), ArrivalTime: mustParseTime("17:30"), Date: mustParseDate("2024-12-15")},
	{Id: 8, From: "Delhi", To: "Mumbai", FlightName: "Air India AI456", Price: "₹8000", Class: Business, DepartureTime: mustParseTime("17:00"), ArrivalTime: mustParseTime("19:30"), Date: mustParseDate("2024-12-15")},
	{Id: 9, From: "Delhi", To: "Mumbai", FlightName: "SpiceJet SG678", Price: "₹5200", Class: Economy, DepartureTime: mustParseTime("18:30"), ArrivalTime: mustParseTime("21:00"), Date: mustParseDate("2024-12-15")},
	{Id: 10, From: "Delhi", To: "Mumbai", FlightName: "Vistara VT789", Price: "₹7500", Class: Premium, DepartureTime: mustParseTime("19:30"), ArrivalTime: mustParseTime("22:00"), Date: mustParseDate("2024-12-15")},

	// Flights on 2024-12-16, same source and destination (Bangalore to Kolkata)
	{Id: 11, From: "Bangalore", To: "Kolkata", FlightName: "Air India AI789", Price: "₹8000", Class: Business, DepartureTime: mustParseTime("05:30"), ArrivalTime: mustParseTime("08:00"), Date: mustParseDate("2024-12-16")},
	{Id: 12, From: "Bangalore", To: "Kolkata", FlightName: "IndiGo 6E456", Price: "₹5500", Class: Economy, DepartureTime: mustParseTime("06:30"), ArrivalTime: mustParseTime("09:00"), Date: mustParseDate("2024-12-16")},
	{Id: 13, From: "Bangalore", To: "Kolkata", FlightName: "SpiceJet SG123", Price: "₹6200", Class: Business, DepartureTime: mustParseTime("07:15"), ArrivalTime: mustParseTime("09:45"), Date: mustParseDate("2024-12-16")},
	{Id: 14, From: "Bangalore", To: "Kolkata", FlightName: "Vistara VT234", Price: "₹4000", Class: Economy, DepartureTime: mustParseTime("09:00"), ArrivalTime: mustParseTime("11:30"), Date: mustParseDate("2024-12-16")},
	{Id: 15, From: "Bangalore", To: "Kolkata", FlightName: "GoAir G8124", Price: "₹7500", Class: Economy, DepartureTime: mustParseTime("10:30"), ArrivalTime: mustParseTime("13:00"), Date: mustParseDate("2024-12-16")},
	{Id: 16, From: "Bangalore", To: "Kolkata", FlightName: "AirAsia AA567", Price: "₹5200", Class: Economy, DepartureTime: mustParseTime("13:30"), ArrivalTime: mustParseTime("16:00"), Date: mustParseDate("2024-12-16")},
	{Id: 17, From: "Bangalore", To: "Kolkata", FlightName: "IndiGo 6E567", Price: "₹6000", Class: Economy, DepartureTime: mustParseTime("15:00"), ArrivalTime: mustParseTime("17:30"), Date: mustParseDate("2024-12-16")},
	{Id: 18, From: "Bangalore", To: "Kolkata", FlightName: "SpiceJet SG890", Price: "₹5800", Class: Premium, DepartureTime: mustParseTime("16:30"), ArrivalTime: mustParseTime("19:00"), Date: mustParseDate("2024-12-16")},
	{Id: 19, From: "Bangalore", To: "Kolkata", FlightName: "Vistara VT123", Price: "₹4500", Class: Economy, DepartureTime: mustParseTime("18:00"), ArrivalTime: mustParseTime("20:30"), Date: mustParseDate("2024-12-16")},
	{Id: 20, From: "Bangalore", To: "Kolkata", FlightName: "IndiGo 6E890", Price: "₹6500", Class: Business, DepartureTime: mustParseTime("19:00"), ArrivalTime: mustParseTime("21:30"), Date: mustParseDate("2024-12-16")},
}

func mustParseTime(value string) time.Time {
	parsedTime, err := time.Parse("15:04", value)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse time: %s", err))
	}
	return parsedTime
}

func mustParseDate(value string) time.Time {
	parsedDate, err := time.Parse("2006-01-02", value)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse date: %s", err))
	}
	return parsedDate
}

var BookFlights []Flight

type Search struct {
	From  string
	To    string
	Date  time.Time
	Class ClassType
}

func UserFlightInput() {
	var from, to, sourceInput, classInput string
	var class ClassType

	fmt.Print("Enter your location:-  ")
	fmt.Scanln(&from)

	fmt.Print("Enter your destination:- ")
	fmt.Scanln(&to)

	fmt.Print("Enter Source Data (format: YYYY-MM-DD):- ")
	fmt.Scanln(&sourceInput)

	const LAYOUT = "2006-01-02"

	date, err := time.Parse(LAYOUT, sourceInput)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}

	if date.Before(time.Now()) {
		println("Invalid Date Input , Don't add past value")
		return
	}

	fmt.Println("Select class (Economy, Premium, Business, First): ")
	fmt.Scanln(&classInput)

	switch classInput {
	case "Economy":
		class = Economy
	case "Premium":
		class = Premium
	case "Business":
		class = Business
	case "First":
		class = First
	default:
		fmt.Println("Invalid class selection. Defaulting to Economy.")
		class = Economy
	}

	// flightmenu.UserMenu()
	fmt.Println("=================================")
	search := Search{
		From:  from,
		To:    to,
		Date:  date,
		Class: class,
	}
	SearchInput(search)

}

func (c ClassType) String() string {
	switch c {
	case Economy:
		return "Economy"
	case Premium:
		return "Premium"
	case Business:
		return "Business"
	case First:
		return "First"
	default:
		return "Unknown"
	}
}

func SearchInput(search Search) {
	var searchResult []Flight

	for _, value := range Flights {
		if search.From == value.From && search.To == value.To && search.Date == value.Date && search.Class == value.Class {
			searchResult = append(searchResult, value)
		}
	}

	if searchResult == nil {
		fmt.Println("No plane available for this date")
		return
	}

	sort.Slice(searchResult, func(i, j int) bool {
		return searchResult[i].Price < searchResult[j].Price
	})

	for _, flight := range searchResult {

		fmt.Printf("FlightId : %d\n", flight.Id)
		fmt.Printf("Flight : %s\n", flight.FlightName)
		fmt.Printf("Class : %s\n", flight.Class.String())
		fmt.Printf("Formatted Departure Time: %s\n", flight.DepartureTime.Format("15:04"))
		fmt.Printf("Formatted Arrival Time: %s\n", flight.ArrivalTime.Format("15:04"))
		fmt.Printf("Formatted Date: %s\n", flight.Date.Format("2006-01-02"))
		fmt.Printf("Price : %s\n", flight.Price)
		fmt.Println("====================================")
	}
}

func AddFlight() {
	var flightName, from, to, sourceInput, departure, arrival, classInput, price string
	var class ClassType

	fmt.Print("Enter Flight Name:-  ")
	fmt.Scanln(&flightName)

	fmt.Print("Enter Source:-  ")
	fmt.Scanln(&from)

	fmt.Print("Enter  destination:- ")
	fmt.Scanln(&to)

	fmt.Print("Enter Source Data (format: YYYY-MM-DD):- ")
	fmt.Scanln(&sourceInput)

	fmt.Print("Enter Departure Time (format: HH:mm):- ")
	fmt.Scanln(&departure)
	fmt.Print("Enter Arrival Time (format: HH:mm):- ")
	fmt.Scanln(&arrival)
	fmt.Print("Enter Price:- ")
	fmt.Scanln(&price)

	const LAYOUT = "2006-01-02"
	date, err := time.Parse(LAYOUT, sourceInput)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		return
	}

	departureTime := mustParseTime(departure)
	arrivalTime := mustParseTime(arrival)

	fmt.Println("Select class (Economy, Premium, Business, First): ")
	fmt.Scanln(&classInput)

	switch classInput {
	case "Economy":
		class = Economy
	case "Premium":
		class = Premium
	case "Business":
		class = Business
	case "First":
		class = First
	default:
		fmt.Println("Invalid class selection. Defaulting to Economy.")
		class = Economy
	}

	newFlight := Flight{
		Id:            len(Flights) + 1,
		FlightName:    flightName,
		From:          from,
		To:            to,
		Price:         price,
		Class:         class,
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
		Date:          date,
	}

	Flights = append(Flights, newFlight)
	fmt.Println("Flight Added Successfully ")

}

func remove(slice []Flight, s int) []Flight {
	return append(slice[:s], slice[s+1:]...)
}

func DeleteFlight() {
	var id, flightIndex int
	fmt.Print("Enter an flight Id : ")
	fmt.Scanln(&id)

	for index, value := range Flights {
		if value.Id == id {
			flightIndex = index
			break
		}
	}

	fmt.Print(Flights[flightIndex])
	Flights = remove(Flights, flightIndex)

	fmt.Print("Deleted Successfully ")
}
