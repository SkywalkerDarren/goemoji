package goemoji

import (
	"fmt"
	"testing"
)

type testItem struct {
	input        string
	replaceEmoji ReplaceFunc
	replaceText  ReplaceFunc
	output       string
	count        int
}

func TestReplaceAllEmoji(t *testing.T) {
	testItems := []testItem{
		{
			input:        "!@#",
			replaceEmoji: func(emoji string) string { return "ASD" },
			replaceText:  func(text string) string { return "QWE" },
			output:       "QWE",
			count:        0,
		},
		{
			input:        "Hello, world! 😄",
			replaceEmoji: func(emoji string) string { return "#" },
			replaceText:  func(text string) string { return text },
			output:       "Hello, world! #",
			count:        1,
		},
		{
			input:        "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿,👨🏼‍🤝‍👨🏿",
			replaceEmoji: func(emoji string) string { return "#" },
			replaceText:  func(text string) string { return text },
			output:       "#(#)(#)(#,#",
			count:        5,
		},
		{
			input:        "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
			replaceEmoji: func(emoji string) string { return "#" },
			replaceText:  func(text string) string { return text },
			output:       "#(#)(#)(##",
			count:        5,
		},
		{
			input:        "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			replaceEmoji: func(emoji string) string { return "#" },
			replaceText:  func(text string) string { return text },
			output:       "#(#)(#)(##)",
			count:        5,
		},
		{
			input:        "(👋👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			replaceEmoji: func(emoji string) string { return "#" },
			replaceText:  func(text string) string { return text },
			output:       "(##)(#)(##)",
			count:        5,
		},
		{
			input:        "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
			replaceEmoji: func(emoji string) string { return fmt.Sprintf("{%s}", emoji) },
			replaceText:  func(text string) string { return fmt.Sprintf("(%s)", text) },
			output:       "(){👨🏼‍🤝‍👨🏿}(){👨🏼‍🤝‍👨🏿}(){👋}(){👨‍👩‍👧‍👧}",
			count:        4,
		},
	}

	for _, item := range testItems {
		s := HandleAll(item.input, item.replaceEmoji, item.replaceText)
		if s != item.output {
			t.Errorf("expected %s, got %s", item.output, s)
		}
		if count := Count(item.input); count != item.count {
			t.Errorf("expected %d, got %d", item.count, count)
		}
	}
}
