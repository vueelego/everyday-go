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

func (n *node) insert(pattern string, parts []string, height int) {
	fmt.Printf("insert pattern=%q, parts=%q height=%q", pattern, parts, height)
	// 退出递归，添加路由height是从0开始；
	// 比如：pattern="/api/user/:id", parts=[api,user,:id],
	// 当len(parts) == height-1时，parts元素都已经添加完了
	// 因此len(parts) == height，已经是没有元素可以insert了，将pattern赋值给最后一个node,退出
	if len(parts) == height {
		n.pattern = pattern
	}

	part := parts[height]       // 获取要添加的节点
	child := n.matchChild(part) // 查找当前节点是否匹配存在，不存在则添加，存在则插入下一个节点，

	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*', // 是否采用模糊匹配或者说精准匹配
		}
	}

}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
