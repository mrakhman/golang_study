package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50

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
	// var conferenceName = "Go Conference"
	// same as ^

	greetUsers()

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookings = bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("There first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is sold out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Printf("We ony have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			}
			// unlike `break` it doesn't stop the loop but goes to another iteration in the loop
			continue
		}
	}

	// city := "London"
	// switch city {
	// case "New Yourk":
	// 	//
	// case "Singapore", "Hong Kong":
	// 	//
	// case " London", "Berlin":
	// 	//
	// case "Mxico City":
	// 	//
	// default:
	// 	fmt.Println("No valid city selected")
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
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
	// asks user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) []UserData {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking so many tickets. You will receive a confirmation email at %v\n", firstName, lastName, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, email)
	fmt.Println("#########")
}
