package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

const MaxWords = 10

type countWords struct {
	word  string
	count int
}

func Top10(text string) []string {
	text = handleText(text)
	if text == "" {
		return nil
	}
	textStruct := strings.Split(text, " ")
	countsWords := getCountWords(textStruct)
	wordsStruct := handleWordsToStruct(countsWords)
	sortWords(wordsStruct)
	if len(wordsStruct) > MaxWords {
		wordsStruct = wordsStruct[:MaxWords]
	}

	return getOnlyWords(wordsStruct)
}

func handleText(text string) string {
	// text = strings.ToLower(text)
	space := regexp.MustCompile(`\s+`)
	text = space.ReplaceAllString(text, " ")
	return text
}

func getCountWords(items []string) map[string]int {
	countWord := make(map[string]int)
	for _, v := range items {
		countWord[v]++
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
		if len(res) == 10 {
			break
		}
	}
	return res
}
