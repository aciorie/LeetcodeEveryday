package trie

// TrieNode 是Trie树的节点
type TrieNode struct {
	children    [26]*TrieNode
	isEndOfWord bool
}

// Trie 是前缀树的根结构
type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, char := range word {
		index := char - '0'
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		node = node.children[index]
		node.isEndOfWord = true
	}
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, char := range word {
		index := char - '0'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, char := range prefix {
		index := char - '0'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}
