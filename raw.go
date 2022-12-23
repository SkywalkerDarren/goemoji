package goemoji

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func newDictTree() *node {
	return &node{NodeMap: make(map[int]*node)}
}

type node struct {
	Char    int
	IsEnd   bool
	NodeMap map[int]*node
}

func (n *node) addNode(code int, isEnd bool) {
	_, ok := n.NodeMap[code]
	if !ok {
		n.NodeMap[code] = &node{Char: code, NodeMap: make(map[int]*node), IsEnd: isEnd}
	}
}

func (n *node) getNode(code int) *node {
	node, ok := n.NodeMap[code]
	if ok {
		return node
	}
	return nil
}

//go:embed official/emoji-sequences.txt
//go:embed official/emoji-zwj-sequences.txt
var fs embed.FS

func readAllEmoji() (*node, error) {
	tree := newDictTree()

	seqFile, err := fs.Open("official/emoji-sequences.txt")
	if seqFile != nil {
		defer seqFile.Close()
	}
	if err != nil {
		return nil, err
	}

	err = convertOfficialEmoji(seqFile, tree)
	if err != nil {
		return nil, err
	}

	zwjFile, err := fs.Open("official/emoji-zwj-sequences.txt")
	if zwjFile != nil {
		defer zwjFile.Close()
	}
	if err != nil {
		return nil, err
	}

	err = convertOfficialEmoji(zwjFile, tree)
	if err != nil {
		return nil, err
	}
	return tree, nil
}

func convertOfficialEmoji(reader io.Reader, tree *node) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] != '#' {
			line = line[:strings.Index(line, ";")]
			line = strings.TrimRight(line, " ")
			codeList := strings.Split(line, " ")
			if len(codeList) > 0 {
				err := handleCodes(codeList, tree)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func handleCodes(codeList []string, root *node) error {
	next := root
	l := len(codeList)
	for i, codeRange := range codeList {
		start, end, found := strings.Cut(codeRange, "..")
		if found {
			if len(codeList) > 1 {
				return fmt.Errorf("code range error: %v", codeRange)
			}
			a, err := strconv.ParseInt(start, 16, 32)
			if err != nil {
				return err
			}
			b, err := strconv.ParseInt(end, 16, 32)
			if err != nil {
				return err
			}
			for i := a; i <= b; i++ {
				next.addNode(int(i), true)
			}
		} else {
			a, err := strconv.ParseInt(start, 16, 32)
			if err != nil {
				return err
			}
			next.addNode(int(a), i == l-1)
			next = next.getNode(int(a))
		}
	}
	return nil
}
