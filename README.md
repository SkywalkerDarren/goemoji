# goemoji😀
![](https://github.com/SkywalkerDarren/goemoji/actions/workflows/go.yml/badge.svg)
![](https://goreportcard.com/badge/github.com/SkywalkerDarren/goemoji)
[![codecov](https://codecov.io/github/SkywalkerDarren/goemoji/branch/master/graph/badge.svg?token=OHSOGISA1M)](https://codecov.io/github/SkywalkerDarren/goemoji)

🚀Fast and ✨simple way to handle text with 👍emoji

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

## Performance

```
go test -bench=. -benchmem -v -run Benchmark ./...
goos: darwin
goarch: arm64
pkg: github.com/SkywalkerDarren/goemoji
BenchmarkRemoveEmojis
BenchmarkRemoveEmojis-8                  9397803               127.2 ns/op            16 B/op          1 allocs/op
BenchmarkRemoveEmojisParallel
BenchmarkRemoveEmojisParallel-8         38755568                30.46 ns/op           16 B/op          1 allocs/op
BenchmarkRemoveText
BenchmarkRemoveText-8                   10777095               111.2 ns/op             0 B/op          0 allocs/op
BenchmarkRemoveTextParallel
BenchmarkRemoveTextParallel-8           51726182                22.37 ns/op            0 B/op          0 allocs/op
BenchmarkReplaceEmojis
BenchmarkReplaceEmojis-8                 9324697               128.1 ns/op            16 B/op          1 allocs/op
BenchmarkReplaceEmojisParallel
BenchmarkReplaceEmojisParallel-8        35846263                29.12 ns/op           16 B/op          1 allocs/op
BenchmarkReplaceText
BenchmarkReplaceText-8                   9069320               131.8 ns/op             8 B/op          1 allocs/op
BenchmarkReplaceTextParallel
BenchmarkReplaceTextParallel-8          41602144                27.75 ns/op            8 B/op          1 allocs/op
BenchmarkReplace
BenchmarkReplace-8                      10482950               114.5 ns/op             0 B/op          0 allocs/op
BenchmarkReplaceParallel
BenchmarkReplaceParallel-8              49429840                22.93 ns/op            0 B/op          0 allocs/op
BenchmarkSplit
BenchmarkSplit-8                         8971143               133.5 ns/op            16 B/op          1 allocs/op
BenchmarkSplitParallel
BenchmarkSplitParallel-8                38264695                30.21 ns/op           16 B/op          1 allocs/op
BenchmarkCount
BenchmarkCount-8                         1587834               755.5 ns/op             0 B/op          0 allocs/op
BenchmarkCountParallel
BenchmarkCountParallel-8                54065012                22.56 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/SkywalkerDarren/goemoji      19.516s
```

## License

MIT License
