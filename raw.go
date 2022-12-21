package goemoji

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func newDictTree() *node {
	return &node{NodeMap: make(map[int64]node)}
}

type node struct {
	Char    int64
	IsEnd   bool
	NodeMap map[int64]node
}

func (n *node) addNode(code int64, isEnd bool) {
	_, ok := n.NodeMap[code]
	if !ok {
		n.NodeMap[code] = node{Char: code, NodeMap: make(map[int64]node), IsEnd: isEnd}
	}
}

func (n *node) getNode(code int64) *node {
	node, ok := n.NodeMap[code]
	if ok {
		return &node
	}
	return nil
}

func readAllEmoji() (*node, error) {
	tree := newDictTree()

	seqFile, err := os.Open("official/emoji-sequences.txt")
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

	zwjFile, err := os.Open("official/emoji-zwj-sequences.txt")
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

func convertOfficialEmoji(file *os.File, tree *node) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] != '#' {
			codeList := strings.Split(strings.TrimSpace(strings.Split(line, ";")[0]), " ")
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
		code := strings.Split(codeRange, "..")
		if len(code) != 1 {
			if len(codeList) > 1 {
				return fmt.Errorf("code range error: %v", codeRange)
			}
			a, err := strconv.ParseInt(code[0], 16, 64)
			if err != nil {
				return err
			}
			b, err := strconv.ParseInt(code[1], 16, 64)
			if err != nil {
				return err
			}
			for i := a; i <= b; i++ {
				next.addNode(i, true)
			}
		} else {
			a, err := strconv.ParseInt(code[0], 16, 64)
			if err != nil {
				return err
			}
			next.addNode(a, i == l-1)
			next = next.getNode(a)
		}
	}
	return nil
}
