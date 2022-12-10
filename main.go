package main

import "fmt"

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

	print(userInfo.FirstName)

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
	fmt.Println("Please enter total number of ticket to book")
	fmt.Scan(&userInfo.NumberofTicketBooked)

	return userInfo

}

func isValidUserData(userData userInfo) bool {

}
