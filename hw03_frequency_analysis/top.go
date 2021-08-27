package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Pair struct {
	Word  string
	Value int
}

var (
	pair []Pair
	str  []string
)

func Top10(text1 string) []string {
	spltdText := strings.Fields(text1)
	sort.Strings(spltdText)
	elements := make(map[string]int)
	for _, value := range spltdText {
		elements[value]++
	}

	for w, v := range elements {
		pair = append(pair, Pair{w, v})
	}

	sort.Slice(pair, func(i, j int) bool {
		if pair[i].Value == pair[j].Value {
			return pair[i].Word < pair[j].Word
		}
		return pair[i].Value > pair[j].Value
	})

	for i, pair := range pair {
		if i < 10 {
			str = append(str, pair.Word)
		}
	}
	return str
}
