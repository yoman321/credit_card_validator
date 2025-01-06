package main

//used to format strings and print to command line
import (
	"bufio"
	"credit_card_validator/users"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addUser(usersSlice []users.User, name string, creditCard int) []users.User {
	return append(usersSlice, users.User{Name: name, CreditCard: creditCard})
}

func isNameInSlice(name string, usersSlice []users.User) (bool, users.User) {
	for _, v := range usersSlice {
		if v.Name == name {
			return true, v
		}
	}
	return false, users.User{}
}

func isCreditCardOnlyDigitsAndRightRange(creditCard string) (bool, int) {
	if len(creditCard) < 13 || len(creditCard) > 16 {
		return false, -1
	}

	creditCardToInt, err := strconv.Atoi(creditCard)
	if err != nil {
		fmt.Println(err)
		return false, -1
	}
	return true, creditCardToInt
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	usersSlice := []users.User{}
	option := ""

	for option != "E" {
		fmt.Println("Please input your option: ")
		fmt.Print("Add User(A), Verify Credit Card(V), Print All Users(P), Exit(E): ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "E":
			os.Exit(0)
		case "A":
			fmt.Print("Please input the name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Please enter a credit card number: ")
			creditCard, _ := reader.ReadString('\n')
			creditCard = strings.TrimSpace(creditCard)
			creditCard = strings.ReplaceAll(creditCard, " ", "")
			isOnlyDigits, digitsOnlyCreditCard := isCreditCardOnlyDigitsAndRightRange(creditCard)

			if !isOnlyDigits {
				fmt.Println("Credit card is not of the right range (not between 13 and 19 digits) and/or does not contain only digits.")
				break
			}

			usersSlice = addUser(usersSlice, name, digitsOnlyCreditCard)

		case "V":
			fmt.Print("Please input your name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			isNameInUsersSlice, user := isNameInSlice(name, usersSlice)
			if !isNameInUsersSlice {
				fmt.Println("Name could not be found.")
				break
			}

			if user.IsCreditCardValid() {
				fmt.Printf("The credit card number %v for %v is valid!\n", name, user.CreditCard)
			} else {
				fmt.Printf("The credit card number %v for %v is NOT valid.\n", name, user.CreditCard)
			}

		case "P":
			for _, v := range usersSlice {
				fmt.Printf("User: %v Credit Card: %v\n", v.Name, v.CreditCard)
			}

		default:
			fmt.Println("Invalid option, please pick from given options")
		}

		fmt.Println()
	}
}
