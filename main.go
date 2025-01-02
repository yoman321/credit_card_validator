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

func addUser(usersSlice []users.User, name string, creditCard string) []users.User {
	creditCardToInt, _ := strconv.Atoi(creditCard)
	return append(usersSlice, users.User{Name: name, CreditCard: creditCardToInt})
}

func isNameInSlice(name string, usersSlice []users.User) (bool, int) {
	for i, v := range usersSlice {
		if v.Name == name {
			return true, i
		}
	}
	return false, -1
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	usersSlice := []users.User{}
	option := ""

	for option != "E" {
		fmt.Println("Please input your option:")
		fmt.Print("Add User(A), Verify Credit Card(V), Exit(E)")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "E":
			break
		case "A":
			fmt.Print("Please input the name: ")
			name, _ := reader.ReadString('\n')
			fmt.Print("Please enter a credit card number: ")
			creditCard, _ := reader.ReadString('\n')

			usersSlice = addUser(usersSlice, name, creditCard)

			fmt.Println(usersSlice) //test
		case "V":
			fmt.Print("Please input your name: ")
			name, _ := reader.ReadString('\n')

			isNameInUsersSlice, nameIndex := isNameInSlice(name, usersSlice)
			if !isNameInUsersSlice {
				break
			}
			if users.IsCreditCardValid(usersSlice[nameIndex].CreditCard) {
				fmt.Printf("The credit card number %v for %v is valid!\n", name, usersSlice[nameIndex].CreditCard)
			} else {
				fmt.Printf("The credit card number %v for %v is not valid.\n", name, usersSlice[nameIndex].CreditCard)
			}

		default:
			fmt.Println("Invalid option, please pick from given options")
		}
	}
}
