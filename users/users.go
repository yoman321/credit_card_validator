package users

type User struct {
	Name       string
	CreditCard int
}

func (user User) IsCreditCardValid() bool {
	creditCard := user.CreditCard
	isEveryTwoDigits := false
	sum := 0

	for creditCard > 0 {
		currDigit := creditCard % 10

		if isEveryTwoDigits {
			sum += currDigit * 2
			isEveryTwoDigits = false
		} else {
			sum += currDigit
			isEveryTwoDigits = true
		}

		creditCard /= 10
	}

	return sum%10 == 0
}
