package trie

type Trie struct {
	root *trieNode
}

type trieNode struct {
	children map[rune]*trieNode
	isAnEnd  bool
}

func NewTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
	}
}
func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for i, char := range word {
		if _, ok := node.children[char]; ok {
			node = node.children[char]
		} else {
			for _, char := range word[i:] {
				node.children[char] = newTrieNode()
				node = node.children[char]
			}
			break
		}
	}
	node.isAnEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		var ok bool
		if node, ok = node.children[char]; !ok {
			return false
		}
	}
	return node.isAnEnd
}

func (t *Trie) Predict(word string) []string {
	var results []string

	//searching part
	node := t.root
	for _, char := range word {
		var ok bool
		if node, ok = node.children[char]; !ok {
			return nil
		}
	}

	addPredictions(&results, node, word)
	return results
}

func addPredictions(results *[]string, node *trieNode, keyWord string) {
	if node.isAnEnd {
		*results = append(*results, keyWord)
	}
	for key, suffixNode := range node.children {
		addPredictions(results, suffixNode, keyWord+string(key))
	}
}
