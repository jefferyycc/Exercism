package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, " ", "")
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < len(isbn); i++ {
		digit := int(isbn[i] - '0')
		if isbn[i] == 'X' && i == 9 {
			digit = 10
		} else if isbn[i] < '0' || isbn[i] > '9' {
			return false
		}
		sum += digit * (10 - i)
	}
	return sum%11 == 0
}
