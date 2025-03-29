# goemoji😀

![](https://github.com/SkywalkerDarren/goemoji/actions/workflows/go.yml/badge.svg)
![](https://goreportcard.com/badge/github.com/SkywalkerDarren/goemoji)
[![codecov](https://codecov.io/github/SkywalkerDarren/goemoji/branch/master/graph/badge.svg?token=OHSOGISA1M)](https://codecov.io/github/SkywalkerDarren/goemoji)

## A High-Performance Emoji Processing Library for Go

goemoji is a **lightweight**, **blazing fast** Go library designed for handling text with emojis. Supporting the latest Unicode emoji standards!

- ✨ **Simple API** - Clean and intuitive interface
- ⚡ **High Performance** - Optimized with parallel processing capabilities
- 🧩 **Feature Rich** - Replace, remove, count, and split operations
- 🔄 **Flexible Processing** - Handle text and emojis separately

## Installation

```bash
go get -u github.com/SkywalkerDarren/goemoji
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/SkywalkerDarren/goemoji"
)

func main() {
	raw := "emoji👋"

	r1 := goemoji.ReplaceEmojis(raw, func(s string) string {
		return "(" + s + ")"
	})
	fmt.Println(r1) // emoji(👋)

	r2 := goemoji.ReplaceText(raw, func(s string) string {
		return "(" + s + ")"
	})
	fmt.Println(r2) // (emoji)👋()

	r3 := goemoji.Replace(raw, func(s string) string {
		return "{" + s + "}"
	}, func(s string) string {
		return "(" + s + ")"
	})
	fmt.Println(r3) // (emoji){👋}()

	r4 := goemoji.RemoveText(raw)
	fmt.Println(r4) // 👋

	r5 := goemoji.RemoveEmojis(raw)
	fmt.Println(r5) // emoji

	c := goemoji.Count(raw)
	fmt.Println(c) // 1

	s := goemoji.Split(raw, true)
	fmt.Println(s) // [emoji 👋]

	goemoji.HandleAll(raw, func(emoji string) {
		// do something with emoji
	}, func(text string) {
		// do something with text
	})
}

```

## Performance Benchmarks

goemoji is optimized for speed with minimal memory allocation:

```
go test -bench=. -benchmem -v -run Benchmark ./...
goos: darwin
goarch: arm64
pkg: github.com/SkywalkerDarren/goemoji
BenchmarkRemoveEmojis-8                  9397803               127.2 ns/op            16 B/op          1 allocs/op
BenchmarkRemoveEmojisParallel-8         38755568                30.46 ns/op           16 B/op          1 allocs/op
BenchmarkRemoveText-8                   10777095               111.2 ns/op             0 B/op          0 allocs/op
BenchmarkRemoveTextParallel-8           51726182                22.37 ns/op            0 B/op          0 allocs/op
BenchmarkReplaceEmojis-8                 9324697               128.1 ns/op            16 B/op          1 allocs/op
BenchmarkReplaceEmojisParallel-8        35846263                29.12 ns/op           16 B/op          1 allocs/op
BenchmarkReplaceText-8                   9069320               131.8 ns/op             8 B/op          1 allocs/op
BenchmarkReplaceTextParallel-8          41602144                27.75 ns/op            8 B/op          1 allocs/op
BenchmarkReplace-8                      10482950               114.5 ns/op             0 B/op          0 allocs/op
BenchmarkReplaceParallel-8              49429840                22.93 ns/op            0 B/op          0 allocs/op
BenchmarkSplit-8                         8971143               133.5 ns/op            16 B/op          1 allocs/op
BenchmarkSplitParallel-8                38264695                30.21 ns/op           16 B/op          1 allocs/op
BenchmarkCount-8                         1587834               755.5 ns/op             0 B/op          0 allocs/op
BenchmarkCountParallel-8                54065012                22.56 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/SkywalkerDarren/goemoji      19.516s
```

## API Reference

- `ReplaceEmojis` - Replace all emojis in text
- `ReplaceText` - Replace all non-emoji content in text
- `Replace` - Replace both text and emojis with different handlers
- `RemoveEmojis` - Remove all emojis from text
- `RemoveText` - Remove all non-emoji content from text
- `Count` - Count the number of emojis in text
- `Split` - Split text into text and emoji segments
- `HandleAll` - Process each text and emoji segment separately

## License

MIT License

## Links

- [GitHub Repository](https://github.com/SkywalkerDarren/goemoji)
- [Go Documentation](https://pkg.go.dev/github.com/SkywalkerDarren/goemoji)
- [Issue Tracker](https://github.com/SkywalkerDarren/goemoji/issues)
