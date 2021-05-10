package trie

type Trie struct {
	root  map[rune]*Trie
	isEnd bool
}

func Construtor() *Trie {
	return &Trie{map[rune]*Trie{}, false}
}

func (tri *Trie) Insert(word string) {
	node := tri
	for _, s := range word {
		if _, ok := node.root[s]; !ok {
			node.root[s] = &Trie{map[rune]*Trie{}, false}
		}
		node = node.root[s]
	}
	node.isEnd = true
}

func (tri *Trie) Search(word string) bool {
	node := tri
	for _, s := range word {
		if _, ok := node.root[s]; !ok {
			return false
		}
		node = node.root[s]
	}
	return node.isEnd
}
