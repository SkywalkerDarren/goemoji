package goemoji

import "strings"

type HandlerFunc = func(s string)
type ReplaceFunc = func(s string) string

var tree *node

func init() {
	root, _ := readAllEmoji()
	tree = root
}

// HandleAll will call the handler function for each emoji and text
func HandleAll(s string, emojiHandler HandlerFunc, textHandler HandlerFunc) {
	next := tree
	startText := 0
	endText := 0
	inEmoji := false
	startEmoji := 0
	endEmoji := 0
	isEmoji := false
	for i, c := range s {
		next = next.getNode(int(c))
		if next != nil {
			if !inEmoji {
				startEmoji = i
			}
			inEmoji = true
			isEmoji = next.IsEnd
		}
		if next == nil && inEmoji {
			inEmoji = false
			if isEmoji {
				endEmoji = i
				text := s[startText:endText]
				textHandler(text)
				emojiHandler(s[startEmoji:endEmoji])
				startText = i
			}
			endText = i
			isEmoji = false
		}
		if next == nil {
			next = tree
			if next.getNode(int(c)) != nil {
				next = next.getNode(int(c))
				if !inEmoji {
					startEmoji = i
				}
				inEmoji = true
				isEmoji = next.IsEnd
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

// ReplaceEmojis will replace all emojis with the replace function
func ReplaceEmojis(s string, replaceFunc ReplaceFunc) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {
		result.WriteString(replaceFunc(emoji))
	}, func(text string) {
		result.WriteString(text)
	})
	return result.String()
}

// ReplaceText will replace all text with the replace function
func ReplaceText(s string, replaceFunc ReplaceFunc) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {
		result.WriteString(emoji)
	}, func(text string) {
		result.WriteString(replaceFunc(text))
	})
	return result.String()
}

// Replace will replace all emojis and text with the replace function
func Replace(s string, replaceEmojiFunc ReplaceFunc, replaceTextFunc ReplaceFunc) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {
		result.WriteString(replaceEmojiFunc(emoji))
	}, func(text string) {
		result.WriteString(replaceTextFunc(text))
	})
	return result.String()
}

// RemoveText will remove all text
func RemoveText(s string) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {
		result.WriteString(emoji)
	}, func(text string) {})
	return result.String()
}

// RemoveEmojis will remove all emojis
func RemoveEmojis(s string) string {
	result := &strings.Builder{}
	HandleAll(s, func(emoji string) {}, func(text string) {
		result.WriteString(text)
	})
	return result.String()
}

// Split will split the string into emojis and text
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

// Count will count the number of emojis
func Count(s string) int {
	i := 0
	HandleAll(s, func(emoji string) {
		i++
	}, func(text string) {})
	return i
}
