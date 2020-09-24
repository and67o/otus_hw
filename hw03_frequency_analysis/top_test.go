package hw03_frequency_analysis //nolint:golint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = true

type test struct {
	text    string
	expected []string
}

func getTestCases() []test {
	return []test{
		{
			text:    text,
			expected: []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"},
		},
		//{
		//	text:    text1,
		//	expected: []string{"ll", "oleg", "e", "p", "o", "sj", "ll!!!!!!", "oleg!", "llloleg", "b"},
		//},
	}
}

func TestTop102(t *testing.T) {
	for _, tst := range getTestCases() {
		fmt.Println(Top10(tst.text))
		require.Subset(t, tst.expected, Top10(tst.text))
	}
}

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			require.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			require.ElementsMatch(t, expected, Top10(text))
		}
	})
}
