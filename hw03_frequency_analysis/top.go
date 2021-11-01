package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var (
	result       []string
	clearSymbols = regexp.MustCompile(`[,.!?-]`)
)

func Top10(text1 string) []string {
	spltdText := strings.Fields(strings.ToLower(text1))
	elements := make(map[string]int)
	for _, value := range spltdText {
		value = clearSymbols.ReplaceAllString(value, "")
		elements[value]++
	}

	type Pair struct {
		Word  string
		Value int
	}

	var pair []Pair

	for w, v := range elements {
		if w != "" {
			pair = append(pair, Pair{w, v})
		}
	}

	sort.Slice(pair, func(i, j int) bool {
		if pair[i].Value == pair[j].Value {
			return pair[i].Word < pair[j].Word
		}
		return pair[i].Value > pair[j].Value
	})

	for i, pair := range pair {
		if i < 10 {
			result = append(result, pair.Word)
		}
	}
	return result
}
