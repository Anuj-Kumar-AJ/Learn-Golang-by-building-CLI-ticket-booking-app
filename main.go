package main

import (
	"fmt"
	"strings"
)

const totalNumberOfTicket = 50

var availableTicket uint = 50

type userInfo struct {
	FirstName            string
	LastName             string
	email                string
	NumberofTicketBooked uint
}

func main() {

	greeting()
	var userInfo = TakeUserInput(availableTicket)

	// user data validation
	checker, result := isValidUserData(userInfo)
	if !checker {
		fmt.Println(result)
		main()
	}

	availableTicket = TicketAvailable(availableTicket, userInfo.NumberofTicketBooked)

	fmt.Printf("Your %v ticket will be send to your email %v", userInfo.NumberofTicketBooked, &userInfo.email)
	main()
}

func greeting() {
	fmt.Printf("hello welcome to the confrence ticket booking  \n")
	fmt.Printf("total number of ticket is  %v and available ticket is %v", totalNumberOfTicket, availableTicket)
}

func TakeUserInput(availableTicket uint) userInfo {

	var userInfo userInfo

	fmt.Println("Please Enter Your detail")
	fmt.Println("Please enter your First Name")
	fmt.Scan(&userInfo.FirstName)

	fmt.Println("Please enter your First Name")
	fmt.Scan(&userInfo.LastName)
	fmt.Println("Please enter your email")
	fmt.Scan(&userInfo.email)
	fmt.Println("Please enter total number of ticket to book")
	fmt.Scan(&userInfo.NumberofTicketBooked)

	return userInfo

}

func isValidUserData(userData userInfo) (bool, string) {

	// email validation
	if !strings.Contains(userData.email, "@") {
		return false, "Invalid Input Please check your emailId\n"
	}

	// checking is ticket available

	if userData.NumberofTicketBooked > availableTicket {
		return false, "No of ticker entered exceeds ticket available\n"
	}
	return true, ""

}

func TicketAvailable(availableTicket uint, NumberofTicketBooked uint) uint {

	return availableTicket - NumberofTicketBooked

}
