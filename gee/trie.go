package gee

import "fmt"

// node trie树节点
type node struct {
	pattern  string  // 待匹配路由（注册路由），例如 /p/:lang
	part     string  // 路由中的一部分（parma），例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 模糊匹配，part 含有 : 或 * 时为true
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}


func ()