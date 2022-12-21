package goemoji

import (
	"strings"
)

type ReplaceFunc = func(emoji string) string

var tree *node

func init() {
	root, _ := readAllEmoji()
	tree = root
}

func HandleAll(s string, emojiHandler ReplaceFunc, textHandler ReplaceFunc) string {

	next := tree
	inEmoji := false
	start := 0
	end := 0
	result := &strings.Builder{}
	sb := &strings.Builder{}
	for i, c := range s {
		next = next.getNode(int64(c))
		if next != nil {
			if !inEmoji {
				start = i
			}
			inEmoji = true
		}
		if next == nil && inEmoji {
			end = i
			inEmoji = false

			result.WriteString(textHandler(sb.String()))
			sb = &strings.Builder{}

			emoji := s[start:end]
			result.WriteString(emojiHandler(emoji))
		}
		if next == nil {
			next = tree
			if next.getNode(int64(c)) != nil {
				next = next.getNode(int64(c))
				if !inEmoji {
					start = i
				}
				inEmoji = true
			}
		}
		if !inEmoji {
			sb.WriteRune(c)
		}
	}
	result.WriteString(textHandler(sb.String()))
	// handle last emoji
	if inEmoji && next != nil && next.IsEnd {
		end = len(s)

		emoji := s[start:end]
		result.WriteString(emojiHandler(emoji))
	}
	return result.String()
}

func Split(s string, withEmoji bool) []string {
	result := make([]string, 0)
	HandleAll(s, func(emoji string) string {
		if withEmoji {
			result = append(result, emoji)
		}
		return ""
	}, func(text string) string {
		result = append(result, text)
		return ""
	})
	return result
}

func Count(s string) int {
	i := 0
	HandleAll(s, func(emoji string) string {
		i++
		return ""
	}, func(text string) string { return "" })
	return i
}
