package goemoji_test

import (
	"testing"

	"github.com/SkywalkerDarren/goemoji"
)

func TestRemoveEmojis(t *testing.T) {
	type testItem struct {
		input  string
		output string
	}
	testItems := []testItem{
		{
			input:  "",
			output: "",
		},
		{
			input:  "!",
			output: "!",
		},
		{
			input:  "!@#",
			output: "!@#",
		},
		{
			input:  "😄",
			output: "",
		},
		{
			input:  "Hello, world! 😄",
			output: "Hello, world! ",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿,👨🏼‍🤝‍👨🏿",
			output: "()()(,",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
			output: "()()(",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "()()()",
		},
		{
			input:  "(👋👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "()()()",
		},
		{
			input:  "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
			output: "",
		},
		{
			input:  "\U0001F96F Hi \U0001F970",
			output: " Hi ",
		},
	}

	for _, item := range testItems {
		s := goemoji.RemoveEmojis(item.input)
		if s != item.output {
			t.Errorf("expected %s, got %s", item.output, s)
		}
	}
}

func BenchmarkRemoveEmojis(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: " Hi ",
	}

	for i := 0; i < b.N; i++ {
		goemoji.RemoveEmojis(item.input)
	}
}

func BenchmarkRemoveEmojisParallel(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: " Hi ",
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goemoji.RemoveEmojis(item.input)
		}
	})
}

func TestRemoveText(t *testing.T) {
	type testItem struct {
		input  string
		output string
	}
	testItems := []testItem{
		{
			input:  "",
			output: "",
		},
		{
			input:  "!",
			output: "",
		},
		{
			input:  "!@#",
			output: "",
		},
		{
			input:  "😄",
			output: "😄",
		},
		{
			input:  "Hello, world! 😄",
			output: "😄",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿,👨🏼‍🤝‍👨🏿",
			output: "👋👨‍👩‍👧‍👧👨‍👩‍👧👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
			output: "👋👨‍👩‍👧‍👧👨‍👩‍👧👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "👋👨‍👩‍👧‍👧👨‍👩‍👧👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
		},
		{
			input:  "(👋👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "👋👨‍👩‍👧‍👧👨‍👩‍👧👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
		},
		{
			input:  "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
			output: "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
		},
		{
			input:  "\U0001F96F Hi \U0001F970",
			output: "🥯🥰",
		},
	}

	for _, item := range testItems {
		s := goemoji.RemoveText(item.input)
		if s != item.output {
			t.Errorf("expected %s, got %s", item.output, s)
		}
	}
}

func BenchmarkRemoveText(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: " Hi ",
	}

	for i := 0; i < b.N; i++ {
		goemoji.RemoveText(item.input)
	}
}

func BenchmarkRemoveTextParallel(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: " Hi ",
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goemoji.RemoveText(item.input)
		}
	})
}

func TestReplaceEmojis(t *testing.T) {
	type testItem struct {
		input  string
		output string
	}
	testItems := []testItem{
		{
			input:  "",
			output: "",
		},
		{
			input:  "!",
			output: "!",
		},
		{
			input:  "!@#",
			output: "!@#",
		},
		{
			input:  "😄",
			output: "#",
		},
		{
			input:  "Hello, world! 😄",
			output: "Hello, world! #",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿,👨🏼‍🤝‍👨🏿",
			output: "#(#)(#)(#,#",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
			output: "#(#)(#)(##",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "#(#)(#)(##)",
		},
		{
			input:  "(👋👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "(##)(#)(##)",
		},
		{
			input:  "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
			output: "####",
		},
		{
			input:  "\U0001F96F Hi \U0001F970",
			output: "# Hi #",
		},
	}

	for _, item := range testItems {
		s := goemoji.ReplaceEmojis(item.input, func(s string) string {
			return "#"
		})
		if s != item.output {
			t.Errorf("expected %s, got %s", item.output, s)
		}
	}
}

func BenchmarkReplaceEmojis(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "# Hi #",
	}

	for i := 0; i < b.N; i++ {
		goemoji.ReplaceEmojis(item.input, func(s string) string {
			return "#"
		})
	}
}

func BenchmarkReplaceEmojisParallel(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "# Hi #",
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goemoji.ReplaceEmojis(item.input, func(s string) string {
				return "#"
			})
		}
	})
}

func TestReplaceText(t *testing.T) {
	type testItem struct {
		input  string
		output string
	}
	testItems := []testItem{
		{
			input:  "",
			output: "#",
		},
		{
			input:  "!",
			output: "#",
		},
		{
			input:  "!@#",
			output: "#",
		},
		{
			input:  "😄",
			output: "#😄#",
		},
		{
			input:  "Hello, world! 😄",
			output: "#😄#",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿,👨🏼‍🤝‍👨🏿",
			output: "#👋#👨‍👩‍👧‍👧#👨‍👩‍👧#👨🏼‍🤝‍👨🏿#👨🏼‍🤝‍👨🏿#",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
			output: "#👋#👨‍👩‍👧‍👧#👨‍👩‍👧#👨🏼‍🤝‍👨🏿#👨🏼‍🤝‍👨🏿#",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "#👋#👨‍👩‍👧‍👧#👨‍👩‍👧#👨🏼‍🤝‍👨🏿#👨🏼‍🤝‍👨🏿#",
		},
		{
			input:  "(👋👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "#👋#👨‍👩‍👧‍👧#👨‍👩‍👧#👨🏼‍🤝‍👨🏿#👨🏼‍🤝‍👨🏿#",
		},
		{
			input:  "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
			output: "#👨🏼‍🤝‍👨🏿#👨🏼‍🤝‍👨🏿#👋#👨‍👩‍👧‍👧#",
		},
		{
			input:  "\U0001F96F Hi \U0001F970",
			output: "#\U0001F96F#\U0001F970#",
		},
	}

	for _, item := range testItems {
		s := goemoji.ReplaceText(item.input, func(s string) string {
			return "#"
		})
		if s != item.output {
			t.Errorf("expected %s, got %s", item.output, s)
		}
	}
}

func BenchmarkReplaceText(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "# Hi #",
	}

	for i := 0; i < b.N; i++ {
		goemoji.ReplaceText(item.input, func(s string) string {
			return "#"
		})
	}
}

func BenchmarkReplaceTextParallel(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "# Hi #",
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goemoji.ReplaceText(item.input, func(s string) string {
				return "#"
			})
		}
	})
}

func TestReplace(t *testing.T) {
	type testItem struct {
		input  string
		output string
	}
	testItems := []testItem{
		{
			input:  "",
			output: "{}",
		},
		{
			input:  "!",
			output: "{!}",
		},
		{
			input:  "!@#",
			output: "{!@#}",
		},
		{
			input:  "😄",
			output: "{}(😄){}",
		},
		{
			input:  "Hello, world! 😄",
			output: "{Hello, world! }(😄){}",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿,👨🏼‍🤝‍👨🏿",
			output: "{}(👋){(}(👨‍👩‍👧‍👧){)(}(👨‍👩‍👧){)(}(👨🏼‍🤝‍👨🏿){,}(👨🏼‍🤝‍👨🏿){}",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
			output: "{}(👋){(}(👨‍👩‍👧‍👧){)(}(👨‍👩‍👧){)(}(👨🏼‍🤝‍👨🏿){}(👨🏼‍🤝‍👨🏿){}",
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "{}(👋){(}(👨‍👩‍👧‍👧){)(}(👨‍👩‍👧){)(}(👨🏼‍🤝‍👨🏿){}(👨🏼‍🤝‍👨🏿){)}",
		},
		{
			input:  "(👋👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: "{(}(👋){}(👨‍👩‍👧‍👧){)(}(👨‍👩‍👧){)(}(👨🏼‍🤝‍👨🏿){}(👨🏼‍🤝‍👨🏿){)}",
		},
		{
			input:  "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
			output: "{}(👨🏼‍🤝‍👨🏿){}(👨🏼‍🤝‍👨🏿){}(👋){}(👨‍👩‍👧‍👧){}",
		},
		{
			input:  "\U0001F96F Hi \U0001F970",
			output: "{}(🥯){ Hi }(🥰){}",
		},
	}

	for _, item := range testItems {
		s := goemoji.Replace(item.input, func(s string) string {
			return "(" + s + ")"
		}, func(s string) string {
			return "{" + s + "}"
		})
		if s != item.output {
			t.Errorf("expected %s, got %s", item.output, s)
		}
	}
}

func BenchmarkReplace(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "",
	}

	for i := 0; i < b.N; i++ {
		goemoji.Replace(item.input, func(s string) string {
			return ""
		}, func(s string) string {
			return ""
		})
	}
}

func BenchmarkReplaceParallel(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "",
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goemoji.Replace(item.input, func(s string) string {
				return ""
			}, func(s string) string {
				return ""
			})
		}
	})
}

func TestSplit(t *testing.T) {
	type testItem struct {
		input      string
		withEmojis bool
		output     []string
	}
	testItems := []testItem{
		{
			input:      "",
			withEmojis: true,
			output:     []string{""},
		},
		{
			input:      "!",
			withEmojis: true,
			output:     []string{"!"},
		},
		{
			input:      "!@#",
			withEmojis: true,
			output:     []string{"!@#"},
		},
		{
			input:      "😄",
			withEmojis: true,
			output:     []string{"", "😄", ""},
		},
		{
			input:      "Hello, world! 😄",
			withEmojis: true,
			output:     []string{"Hello, world! ", "😄", ""},
		},
		{
			input:      "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿,👨🏼‍🤝‍👨🏿",
			withEmojis: true,
			output:     []string{"", "👋", "(", "👨‍👩‍👧‍👧", ")(", "👨‍👩‍👧", ")(", "👨🏼‍🤝‍👨🏿", ",", "👨🏼‍🤝‍👨🏿", ""},
		},
		{
			input:      "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
			withEmojis: true,
			output:     []string{"", "👋", "(", "👨‍👩‍👧‍👧", ")(", "👨‍👩‍👧", ")(", "👨🏼‍🤝‍👨🏿", "", "👨🏼‍🤝‍👨🏿", ""},
		},
		{
			input:      "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			withEmojis: true,
			output:     []string{"", "👋", "(", "👨‍👩‍👧‍👧", ")(", "👨‍👩‍👧", ")(", "👨🏼‍🤝‍👨🏿", "", "👨🏼‍🤝‍👨🏿", ")"},
		},
		{
			input:      "(👋👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			withEmojis: true,
			output:     []string{"(", "👋", "", "👨‍👩‍👧‍👧", ")(", "👨‍👩‍👧", ")(", "👨🏼‍🤝‍👨🏿", "", "👨🏼‍🤝‍👨🏿", ")"},
		},
		{
			input:      "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
			withEmojis: true,
			output:     []string{"", "👨🏼‍🤝‍👨🏿", "", "👨🏼‍🤝‍👨🏿", "", "👋", "", "👨‍👩‍👧‍👧", ""},
		},
		{
			input:      "\U0001F96F Hi \U0001F970",
			withEmojis: true,
			output:     []string{"", "🥯", " Hi ", "🥰", ""},
		},
		{
			input:      "",
			withEmojis: false,
			output:     []string{""},
		},
		{
			input:      "!",
			withEmojis: false,
			output:     []string{"!"},
		},
		{
			input:      "!@#",
			withEmojis: false,
			output:     []string{"!@#"},
		},
		{
			input:      "😄",
			withEmojis: false,
			output:     []string{"", ""},
		},
		{
			input:      "Hello, world! 😄",
			withEmojis: false,
			output:     []string{"Hello, world! ", ""},
		},
		{
			input:      "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿,👨🏼‍🤝‍👨🏿",
			withEmojis: false,
			output:     []string{"", "(", ")(", ")(", ",", ""},
		},
		{
			input:      "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
			withEmojis: false,
			output:     []string{"", "(", ")(", ")(", "", ""},
		},
		{
			input:      "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			withEmojis: false,
			output:     []string{"", "(", ")(", ")(", "", ")"},
		},
		{
			input:      "(👋👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			withEmojis: false,
			output:     []string{"(", "", ")(", ")(", "", ")"},
		},
		{
			input:      "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
			withEmojis: false,
			output:     []string{"", "", "", "", ""},
		},
		{
			input:      "\U0001F96F Hi \U0001F970",
			withEmojis: false,
			output:     []string{"", " Hi ", ""},
		},
	}

	for _, item := range testItems {
		s := goemoji.Split(item.input, item.withEmojis)
		if len(s) != len(item.output) {
			t.Errorf("input: %s, expected %+v %d, got %+v %d", item.input, item.output, len(item.output), s, len(s))
			continue
		}
		for i := range s {
			if s[i] != item.output[i] {
				t.Errorf("input: %s, expected %+v, got %+v", item.input, item.output, s)
				break
			}
		}
	}
}

func BenchmarkSplit(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "",
	}

	for i := 0; i < b.N; i++ {
		goemoji.Split(item.input, true)
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "",
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goemoji.Split(item.input, true)
		}
	})
}

func TestCount(t *testing.T) {
	type testItem struct {
		input  string
		output int
	}
	testItems := []testItem{
		{
			input:  "",
			output: 0,
		},
		{
			input:  "!",
			output: 0,
		},
		{
			input:  "!@#",
			output: 0,
		},
		{
			input:  "😄",
			output: 1,
		},
		{
			input:  "Hello, world! 😄",
			output: 1,
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿,👨🏼‍🤝‍👨🏿",
			output: 5,
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿",
			output: 5,
		},
		{
			input:  "👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: 5,
		},
		{
			input:  "(👋👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿)",
			output: 5,
		},
		{
			input:  "👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿👋👨‍👩‍👧‍👧",
			output: 4,
		},
		{
			input:  "\U0001F96F Hi \U0001F970",
			output: 2,
		},
	}

	for _, item := range testItems {
		s := goemoji.Count(item.input)
		if s != item.output {
			t.Errorf("expected %d, got %d", item.output, s)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "",
	}

	for i := 0; i < b.N; i++ {
		goemoji.Count(item.input)
	}
}

func BenchmarkCountParallel(b *testing.B) {
	type testItem struct {
		input  string
		output string
	}
	item := testItem{
		input:  "說 Hi 殺",
		output: "",
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			goemoji.Count(item.input)
		}
	})
}
