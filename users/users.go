package users

import "fmt"

type User struct {
	Name       string
	CreditCard int
}

func IsCreditCardValid(creditCard int) bool {
	fmt.Printf("For now credit card number %v is valid -> need to implement actual verification afterward", creditCard)
	return true
}
