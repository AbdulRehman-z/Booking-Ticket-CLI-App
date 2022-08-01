package main

// fmt is the formate package for go and has several functionalites
import (
	"fmt"
	"sync"
	"time"
)

const confTickets = 50

var confName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)
var waitGroup = sync.WaitGroup{}

type userData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

func main() {
	// method (1) greet users on startup
	greetUsers()

	firstName, lastName, email, userTickets := getUserInputs()

	// method (2) user inputs validations
	isValidName, isValidEmail, userTicketsValidation := validateUserInputs(firstName, lastName, email, userTickets, remainingTickets)
	// Gaurd Clause
	if isValidEmail && isValidName && userTicketsValidation {

		bookTicket(userTickets, firstName, lastName, email)
		waitGroup.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		fmt.Printf("Bookings: %v\n", bookings)
		// method (3)  Get return value from getFirstNames() method
		firstNames := getFirstNames()
		fmt.Printf("Bookers: %v\n", firstNames)

		// Gaurd Clause
		if remainingTickets == 0 {
			fmt.Printf("All tickets sold out! \n")
		}
	} else {
		if !isValidName {
			fmt.Printf("Invalid firstname or lastname\n")
		}
		if !isValidEmail {
			fmt.Printf("Invalid email address! @ sign missing\n")
		}
		if !userTicketsValidation {
			fmt.Printf("Number of ticket you entered is not valid\n")
		}

	}
	waitGroup.Wait()
}

//////////////////////////////////
////Greet User On Startup////

func greetUsers() {

	fmt.Printf("Welcome to %v\n", confName)
	fmt.Printf("We have total of %v tickets available and still %v left\n", confTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend\n")
}

////////////////////////////////////////////////////////////////
////Print Firsts Names of all the users that booked ticket////

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

//////////////////////////////////
////User inputs validations////

func getUserInputs() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Please enter your firstname:\n")
	fmt.Scan(&firstName)

	fmt.Printf("Please enter your lastname:\n")
	fmt.Scan(&lastName)

	fmt.Printf("Please enter your email address:\n")
	fmt.Scan(&email)

	fmt.Printf("How much tickets you want to get:\n")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Username %v %v, email address %v booked %v tickets! We will soon send you a conformation email at %v \n", firstName, lastName, email, userTickets, email)
	fmt.Printf("remainingTickets: %v\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket: %v to email address: %v", ticket, email)
	fmt.Println("########################")
	waitGroup.Done()
}
