package main

import (
	"fmt"
	"net"
	"net/mail"
	"strings"
)

// isValidEmail checks if the email format is valid.
func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// isDomainValid checks if the email domain has valid MX records.
func isDomainValid(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]

	mxRecords, err := net.LookupMX(domain)
	fmt.Println("mx", mxRecords, domain)
	return err == nil && len(mxRecords) > 0
}

// func main() {
// 	emails := []string{
// 		"valid@example.com",
// 		"invalid-email",
// 		"another.valid.email@example.co.uk",
// 		"@missinglocalpart.com",
// 		"prassna44@gmail.com",
// 	}

// 	for _, email := range emails {
// 		if isValidEmail(email) && isDomainValid(email) {
// 			fmt.Printf("%s is valid and has a valid domain.\n", email)
// 		} else {
// 			fmt.Printf("%s is invalid or has an invalid domain.\n", email)
// 		}
// 	}
// }
