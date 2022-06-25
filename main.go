package main

import (
	"fmt"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is fully booked!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // & - Go pointer

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName,
		lastName,
		email,
		userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v ticket remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("###############")
}
