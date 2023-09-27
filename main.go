package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your First Name: ")
		_, _ = fmt.Scan(&firstName)

		fmt.Println("Enter your Last Name: ")
		_, _ = fmt.Scan(&lastName)

		fmt.Println("Enter your Email: ")
		_, _ = fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		_, _ = fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(
			firstName,
			lastName,
			email,
			userTickets,
			remainingTickets)

		if isValidTicketNumber && isValidName && isValidEmail {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			//fmt.Printf("The Whole Slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("Slice type: %T\n", bookings)
			//fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
				firstName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next")
				break
			}
		} else {
			//fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n",
			//	remainingTickets, userTickets)
			if !isValidName {
				fmt.Println("First Name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}

}

func validateUserInput(firstName string,
	lastName string,
	email string,
	userTickets uint,
	remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
