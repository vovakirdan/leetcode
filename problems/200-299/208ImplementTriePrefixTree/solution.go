package trieprefixtree

// Trie represents a trie data structure
type Trie struct {
	Children   map[rune]*Trie
	IsTerminal bool
}

// Constructor creates a new Trie
func Constructor() Trie {
	return Trie{
		Children: make(map[rune]*Trie),
	}
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		if node.Children[c] == nil {
			node.Children[c] = &Trie{
				Children: make(map[rune]*Trie),
			}
		}
		node = node.Children[c]
	}
	node.IsTerminal = true
}

// Search returns true if the word is in the trie
func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word {
		exists, ok := node.Children[c]
		if !ok {
			return false
		}
		node = exists
	}
	return node.IsTerminal
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, c := range prefix {
		exists, ok := node.Children[c]
		if !ok {
			return false
		}
		node = exists
	}
	return true
}
