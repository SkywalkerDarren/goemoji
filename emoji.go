package goemoji

import "strings"

type HandlerFunc = func(s string)
type ReplaceFunc = func(s string) string

var tree *node

func init() {
	root, _ := readAllEmoji()
	tree = root
}

func HandleAll(s string, emojiHandler HandlerFunc, textHandler HandlerFunc) {
	next := tree
	startText := 0
	endText := 0
	inEmoji := false
	startEmoji := 0
	endEmoji := 0
	for i, c := range s {
		next = next.getNode(int(c))
		if next != nil {
			if !inEmoji {
				startEmoji = i
			}
			inEmoji = true
		}
		if next == nil && inEmoji {
			endEmoji = i
			inEmoji = false
			text := s[startText:endText]
			textHandler(text)
			emojiHandler(s[startEmoji:endEmoji])
			startText = i
			endText = i
		}
		if next == nil {
			next = tree
			if next.getNode(int(c)) != nil {
				next = next.getNode(int(c))
				if !inEmoji {
					startEmoji = i
				}
				inEmoji = true
			}
		}
		if !inEmoji {
			endText = i + 1
		}
	}
	// handle last emoji
	if inEmoji && next != nil && next.IsEnd {
		textHandler(s[startText:endText])
		endEmoji = len(s)
		emojiHandler(s[startEmoji:endEmoji])
		textHandler(s[endEmoji:])
	} else {
		textHandler(s[startText:])
	}
}

func ReplaceEmojis(s string, replaceFunc ReplaceFunc) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {
		result.WriteString(replaceFunc(emoji))
	}, func(text string) {
		result.WriteString(text)
	})
	return result.String()
}

func ReplaceText(s string, replaceFunc ReplaceFunc) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {
		result.WriteString(emoji)
	}, func(text string) {
		result.WriteString(replaceFunc(text))
	})
	return result.String()
}

func Replace(s string, replaceEmojiFunc ReplaceFunc, replaceTextFunc ReplaceFunc) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {
		result.WriteString(replaceEmojiFunc(emoji))
	}, func(text string) {
		result.WriteString(replaceTextFunc(text))
	})
	return result.String()
}

func RemoveText(s string) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {
		result.WriteString(emoji)
	}, func(text string) {})
	return result.String()
}

func RemoveEmojis(s string) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {}, func(text string) {
		result.WriteString(text)
	})
	return result.String()
}

func Split(s string, withEmoji bool) []string {
	result := make([]string, 0)
	HandleAll(s, func(emoji string) {
		if withEmoji {
			result = append(result, emoji)
		}
	}, func(text string) {
		result = append(result, text)
	})
	return result
}

func Count(s string) int {
	i := 0
	HandleAll(s, func(emoji string) {
		i++
	}, func(text string) {})
	return i
}
