package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)



func Top10(text string) []string {
	text = handleText(text)
	textStruct := strings.Split(text, " ")
	countsWords := countSame(textStruct)
	sort.Slice(countsWords, func(i, j int) bool {
		fmt.Println(countsWords[i], countsWords[j])
		return true
	})

	//for k, v := range countsWords {
	//	fmt.Printf("Item : %s --- Count : %d\n", k, v)
	//}
	return nil
}

func handleText(text string) string {
	text = strings.ToLower(text)
	space := regexp.MustCompile("\\s+")
	text = space.ReplaceAllString(text, " ")
	return text
}

func countSame(items []string) map[string]int {
	countSame := make(map[string]int)
	for _, item := range items {
		_, ok := countSame[item]

		if ok {
			countSame[item] += 1
		} else {
			countSame[item] = 1
		}
	}
	return countSame
}
