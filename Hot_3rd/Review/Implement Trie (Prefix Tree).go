package review

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}
type Trie struct {
	root *TrieNode
}

func Constructor_208() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, char := range word {
		idx := char - 'a' // 计算字符在数组中的索引
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{} // 如果子节点不存在，则创建
		}
		node = node.children[idx] // 移动到下一个节点
	}
	node.isEnd = true // 标记单词的最后一个字符所在的节点为单词结尾
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, char := range word {
		idx := char - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, char := range prefix {
		idx := char - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return true
}
