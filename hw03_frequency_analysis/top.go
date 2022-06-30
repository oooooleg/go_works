package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const (
	punctuationCutset = "?!.,:;-"
	topCount          = 10
)

func Top10(text string) (result []string) {
	if text == "" {
		return
	}

	split := strings.Fields(text)
	frequencies := map[string]int{}

	for _, word := range split {
		cleanWord := strings.ToLower(strings.Trim(word, punctuationCutset))
		if len(cleanWord) == 0 {
			continue
		}
		frequencies[cleanWord]++
	}

	type kv struct {
		Key   string
		Value int
	}

	sortedFreqs := []kv{}
	for k, v := range frequencies {
		sortedFreqs = append(sortedFreqs, kv{k, v})
	}

	sort.Slice(sortedFreqs, func(i, j int) bool {
		if sortedFreqs[i].Value == sortedFreqs[j].Value {
			return sortedFreqs[i].Key < sortedFreqs[j].Key
		}
		return sortedFreqs[i].Value > sortedFreqs[j].Value
	})

	for i, kv := range sortedFreqs {
		if i == topCount {
			break
		}
		result = append(result, kv.Key)
	}

	return
}
