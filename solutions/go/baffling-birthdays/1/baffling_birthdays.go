package bafflingbirthdays

import (
	rand "math/rand/v2"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	seen := make(map[uint16]struct{}, len(dates))
	for _, date := range dates {
		birthday := uint16(date.Month())*32 + uint16(date.Day())
		if _, exists := seen[birthday]; exists {
			return true
		}
		seen[birthday] = struct{}{}
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	birthdates := make([]time.Time, size)
	base := time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC)
	for i := range birthdates {
		birthdates[i] = base.AddDate(0, 0, rand.IntN(365))
	}
	return birthdates
}

func EstimatedProbability(size int) float64 {
	const rounds = 1024 * 8
	shared := 0
	for range rounds {
		if SharedBirthday(RandomBirthdates(size)) {
			shared++
		}
	}
	return 100 * float64(shared) / rounds
}
