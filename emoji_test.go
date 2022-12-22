package goemoji_test

import (
	"fmt"
	"goemoji"
	"testing"
)

type testItem struct {
	input        string
	replaceEmoji goemoji.ReplaceFunc
	replaceText  goemoji.ReplaceFunc
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
		{
			input:        "\U0001F96F Hi \U0001F970",
			replaceEmoji: func(emoji string) string { return "" },
			replaceText:  func(text string) string { return text },
			output:       " Hi ",
			count:        2,
		},
	}

	for _, item := range testItems {
		s := goemoji.HandleAll(item.input, item.replaceEmoji, item.replaceText)
		if s != item.output {
			t.Errorf("expected %s, got %s", item.output, s)
		}
		if count := goemoji.Count(item.input); count != item.count {
			t.Errorf("expected %d, got %d", item.count, count)
		}
	}
}

func BenchmarkEmoji(b *testing.B) {
	item := testItem{
		input:        "說 Hi 殺",
		replaceEmoji: func(emoji string) string { return "" },
		replaceText:  func(text string) string { return text },
	}

	for i := 0; i < b.N; i++ {
		goemoji.HandleAll(item.input, item.replaceEmoji, item.replaceText)
	}
}

func BenchmarkEmojiParallel(b *testing.B) {
	item := testItem{
		input:        "說 Hi 殺",
		replaceEmoji: func(emoji string) string { return "" },
		replaceText:  func(text string) string { return text },
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goemoji.HandleAll(item.input, item.replaceEmoji, item.replaceText)
		}
	})
}
