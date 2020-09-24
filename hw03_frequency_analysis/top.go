package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

const MaxWords = 10

var regEx = `\s+`

type countWords struct {
	word  string
	count int
}

func Top10(text string) []string {
	if text == "" {
		return nil
	}

	textStruct := getStructWords(text)
	countsWords := getCountWords(textStruct)
	wordsStruct := handleWordsToStruct(countsWords)
	sortWords(wordsStruct)

	return getOnlyWords(wordsStruct)
}

func getStructWords(text string) []string {
	text = strings.ToLower(text)
	space := regexp.MustCompile(regEx)
	text = space.ReplaceAllString(text, " ")

	return strings.Split(text, " ")
}

func getCountWords(items []string) map[string]int {
	countWord := make(map[string]int)
	for _, word := range items {
		if word == "" || word == "-" {
			continue
		}
		countWord[word]++
	}
	return countWord
}

func handleWordsToStruct(items map[string]int) []countWords {
	wordsStruct := make([]countWords, 0, len(items))
	for count, word := range items {
		wordsStruct = append(wordsStruct, countWords{count, word})
	}
	return wordsStruct
}

func sortWords(wordsStruct []countWords) {
	sort.Slice(wordsStruct, func(i, j int) bool {
		return wordsStruct[i].count > wordsStruct[j].count
	})
}

func getOnlyWords(wordsStruct []countWords) []string {
	res := make([]string, 0, len(wordsStruct))
	for _, word := range wordsStruct {
		res = append(res, word.word)
		if len(res) == MaxWords {
			break
		}
	}

	if len(wordsStruct) > MaxWords {
		res = res[:MaxWords]
	}

	return res
}
