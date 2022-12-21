# goemoji
Simple way to handle text with emoji

## Usage

```go
s := goemoji.HandleAll("👋(👨‍👩‍👧‍👧)(👨‍👩‍👧)(👨🏼‍🤝‍👨🏿👨🏼‍🤝‍👨🏿", func(emoji string) string { return "#" }, func(text string) string { return text })
expected := "#(#)(#)(##"
s == expected // true
```
