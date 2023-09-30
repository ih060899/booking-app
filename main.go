package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(
			firstName,
			lastName,
			email,
			userTickets,
			remainingTickets)

		if isValidTicketNumber && isValidName && isValidEmail {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next")
				break
			}
		} else {
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

func bookTicket(userTickets uint,
	firstName string,
	lastName string,
	email string) {
	remainingTickets = remainingTickets - userTickets
	// create a map for user
	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTickets,
	}

	//userData := make(map[string]string)
	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", userData)

	//fmt.Printf("The Whole Slice: %v\n", bookings)
	//fmt.Printf("The first value: %v\n", bookings[0])
	//fmt.Printf("Slice type: %T\n", bookings)
	//fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, email)
	fmt.Println("#######################")
}
