package ledger

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

const (
	dateWidth        = 10
	descriptionWidth = 25
	changeWidth      = 13
)

func FormatLedger(currency, locale string, entries []Entry) (string, error) {
	if currency != "USD" && currency != "EUR" {
		return "", fmt.Errorf("unsupported currency: %q", currency)
	}
	if locale != "en-US" && locale != "nl-NL" {
		return "", fmt.Errorf("unsupported locale: %q", locale)
	}

	orderedEntries := append([]Entry(nil), entries...)
	for _, entry := range orderedEntries {
		if _, err := time.Parse("2006-01-02", entry.Date); err != nil {
			return "", fmt.Errorf("invalid date %q: %w", entry.Date, err)
		}
	}
	sort.Slice(orderedEntries, func(i, j int) bool {
		left, right := orderedEntries[i], orderedEntries[j]
		if left.Date != right.Date {
			return left.Date < right.Date
		}
		if left.Description != right.Description {
			return left.Description < right.Description
		}
		return left.Change < right.Change
	})

	var output strings.Builder
	output.WriteString(header(locale))
	for _, entry := range orderedEntries {
		date, _ := time.Parse("2006-01-02", entry.Date)
		description := truncate(entry.Description, descriptionWidth)
		amount := formatAmount(entry.Change, currency, locale)
		fmt.Fprintf(&output, "%-*s | %-*s | %*s\n",
			dateWidth, formatDate(date, locale),
			descriptionWidth, description,
			changeWidth, amount,
		)
	}
	return output.String(), nil
}

func header(locale string) string {
	dateLabel, descriptionLabel, changeLabel := "Date", "Description", "Change"
	if locale == "nl-NL" {
		dateLabel, descriptionLabel, changeLabel = "Datum", "Omschrijving", "Verandering"
	}
	return fmt.Sprintf("%-*s | %-*s | %-*s\n",
		dateWidth, dateLabel, descriptionWidth, descriptionLabel, changeWidth, changeLabel)
}

func formatDate(date time.Time, locale string) string {
	if locale == "nl-NL" {
		return date.Format("02-01-2006")
	}
	return date.Format("01/02/2006")
}

func truncate(value string, maxRunes int) string {
	runes := []rune(value)
	if len(runes) > maxRunes {
		return string(runes[:maxRunes-3]) + "..."
	}
	return value
}

func formatAmount(change int, currency, locale string) string {
	amount := int64(change)
	negative := amount < 0
	if negative {
		amount = -amount
	}

	major, minor := amount/100, amount%100
	separator := ","
	decimal := "."
	if locale == "nl-NL" {
		separator = "."
		decimal = ","
	}

	symbol := "$"
	if currency == "EUR" {
		symbol = "€"
	}
	majorText := groupThousands(major, separator)

	if locale == "nl-NL" {
		sign := ""
		if negative {
			sign = "-"
		}
		return fmt.Sprintf("%s %s%s%s%02d ", symbol, sign, majorText, decimal, minor)
	}

	value := fmt.Sprintf("%s%s%s%02d", symbol, majorText, decimal, minor)
	if negative {
		return "(" + value + ")"
	}
	return value + " "
}

func groupThousands(value int64, separator string) string {
	digits := fmt.Sprintf("%d", value)
	firstGroupLength := len(digits) % 3
	if firstGroupLength == 0 {
		firstGroupLength = 3
	}

	var grouped strings.Builder
	grouped.WriteString(digits[:firstGroupLength])
	for i := firstGroupLength; i < len(digits); i += 3 {
		grouped.WriteString(separator)
		grouped.WriteString(digits[i : i+3])
	}
	return grouped.String()
}
