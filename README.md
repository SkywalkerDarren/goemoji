# goemoji
🚀Fast and ✨simple way to handle text with emoji

## Installing

```bash
$ go get -u github.com/SkywalkerDarren/goemoji
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/SkywalkerDarren/goemoji"
)

func main() {
	s := goemoji.HandleAll("👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿", func(emoji string) string { return "#" }, func(text string) string { return text })
	expected := "#(#)(#)(##"
	fmt.Println(s == expected) // true
}
```

## Performance

```
go test -bench=. -benchmem -v -run Benchmark ./...
goos: darwin
goarch: arm64
pkg: goemoji
BenchmarkEmoji
BenchmarkEmoji-8                 3528232               310.1 ns/op           200 B/op         10 allocs/op
BenchmarkEmojiParallel
BenchmarkEmojiParallel-8        10200574               120.9 ns/op           200 B/op         10 allocs/op
PASS
ok      goemoji 3.157s
```

## License

MIT License
