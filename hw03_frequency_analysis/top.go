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
		wordsStruct = wordsStruct[0:MaxWords]
	}

	return getOnlyWords(wordsStruct)
}

func getOnlyWords(wordsStruct []countWords) []string {
	res := make([]string, len(wordsStruct))
	for _, v := range wordsStruct {
		res = append(res, v.word)
		if len(res) == 10 {
			break
		}
	}
	return res
}

func sortWords(wordsStruct []countWords) {
	sort.Slice(wordsStruct, func(i, j int) bool {
		return wordsStruct[i].count > wordsStruct[j].count
	})
}

func handleText(text string) string {
	// text = strings.ToLower(text)
	space := regexp.MustCompile(`\s+`)
	text = space.ReplaceAllString(text, " ")
	return text
}

func handleWordsToStruct(items map[string]int) []countWords {
	wordsStruct := make([]countWords, len(items))
	for i, v := range items {
		wordsStruct = append(wordsStruct, countWords{i, v})
	}
	return wordsStruct
}

func getCountWords(items []string) map[string]int {
	countWord := make(map[string]int)
	for _, v := range items {
		countWord[v]++
	}
	return countWord
}
