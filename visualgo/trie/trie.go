package trie

import "strings"

// Node 定义节点，存储数据
type Node struct {
	children map[rune]*Node
	isTail   bool // 是否是最后一个节点
}

func NewNode() *Node {
	return &Node{children: make(map[rune]*Node)}
}

// Trie 前缀数结构
type Trie struct {
	root *Node
}

func New() *Trie {
	trie := &Trie{
		root: NewNode(),
	}

	return trie
}

// Insert 插入词语
func (t *Trie) Insert(word string) {
	word = strings.ReplaceAll(word, " ", "")
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewNode()
		}
		node = node.children[ch]
	}
	node.isTail = true
}

// Search 查询
func (t *Trie) Search(word string) bool {
	word = strings.ReplaceAll(word, " ", "")
	node := t.findNode(word)
	return node != nil && node.isTail
}

func (t *Trie) findNode(word string) *Node {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return nil
		}
		node = node.children[ch]
	}
	return node
}
