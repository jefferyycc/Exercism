package anagram

import (
	"maps"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	lower := strings.ToLower(subject)
	// key := sortedLetters(lower)
	key2 := letterCounts(lower)
	var result []string
	for _, c := range candidates {
		lc := strings.ToLower(c)
		if lc == lower {
			continue
		}
		// if sortedLetters(lc) == key{
		if maps.Equal(letterCounts(lc), key2) {
			result = append(result, c)
		}
	}
	return result
}

// sorting: O(N log N)
// func sortedLetters(s string) string {
// 	r := []rune(s)
// 	slices.Sort(r)
// 	return string(r)
// }

// count letter: O(N)
func letterCounts(s string) map[rune]int {
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	return count
}
