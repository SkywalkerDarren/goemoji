# goemoji😀
![](https://github.com/SkywalkerDarren/goemoji/actions/workflows/go.yml/badge.svg)
![](https://goreportcard.com/badge/github.com/SkywalkerDarren/goemoji)
[![codecov](https://codecov.io/github/SkywalkerDarren/goemoji/branch/master/graph/badge.svg?token=OHSOGISA1M)](https://codecov.io/github/SkywalkerDarren/goemoji)

🚀Fast and ✨simple way to handle text with 👍emoji
Support unicode emoji 15.0

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
BenchmarkRemoveEmojis-8                  4751536               244.5 ns/op           160 B/op          7 allocs/op
BenchmarkRemoveEmojisParallel
BenchmarkRemoveEmojisParallel-8         12941796                93.93 ns/op          160 B/op          7 allocs/op
BenchmarkRemoveText
BenchmarkRemoveText-8                    5370500               219.4 ns/op           144 B/op          6 allocs/op
BenchmarkRemoveTextParallel
BenchmarkRemoveTextParallel-8           14664808                81.95 ns/op          144 B/op          6 allocs/op
BenchmarkReplaceEmojis
BenchmarkReplaceEmojis-8                 4845986               252.3 ns/op           160 B/op          7 allocs/op
BenchmarkReplaceEmojisParallel
BenchmarkReplaceEmojisParallel-8        12617378                92.83 ns/op          160 B/op          7 allocs/op
BenchmarkReplaceText
BenchmarkReplaceText-8                   4910674               242.0 ns/op           152 B/op          7 allocs/op
BenchmarkReplaceTextParallel
BenchmarkReplaceTextParallel-8          12957044               101.2 ns/op           152 B/op          7 allocs/op
BenchmarkReplace
BenchmarkReplace-8                       4801255               231.6 ns/op           144 B/op          6 allocs/op
BenchmarkReplaceParallel
BenchmarkReplaceParallel-8              13167669               105.1 ns/op           144 B/op          6 allocs/op
BenchmarkSplit
BenchmarkSplit-8                         4740147               254.9 ns/op           160 B/op          7 allocs/op
BenchmarkSplitParallel
BenchmarkSplitParallel-8                12324456                93.58 ns/op          160 B/op          7 allocs/op
BenchmarkCount
BenchmarkCount-8                         5459313               219.7 ns/op           144 B/op          6 allocs/op
BenchmarkCountParallel
BenchmarkCountParallel-8                14243266                83.16 ns/op          144 B/op          6 allocs/op
PASS
ok      github.com/SkywalkerDarren/goemoji      19.363s
```

## License

MIT License
