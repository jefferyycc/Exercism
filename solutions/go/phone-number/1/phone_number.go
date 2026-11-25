package phonenumber

import (
	"errors"
)

var errInvalidPhoneNumber = errors.New("invalid phone number")

func digits(phoneNumber string) (string, error) {
	var number string

	for i := 0; i < len(phoneNumber); i++ {
		char := phoneNumber[i]
		switch {
		case char >= '0' && char <= '9':
			number += string(char)
		case char == ' ' || char == '-' || char == '.' || char == '(' || char == ')':
			// Ignore common phone number formatting characters.
		case char == '+' && i == 0:
			// A leading plus is allowed for the country code.
		default:
			return "", errInvalidPhoneNumber
		}
	}

	if len(number) == 11 && number[0] == '1' {
		number = number[1:]
	}
	if len(number) != 10 || number[0] < '2' || number[0] > '9' || number[3] < '2' || number[3] > '9' {
		return "", errInvalidPhoneNumber
	}

	return number, nil
}

func Number(phoneNumber string) (string, error) {
	return digits(phoneNumber)
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := digits(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := digits(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + number[:3] + ") " + number[3:6] + "-" + number[6:], nil
}
