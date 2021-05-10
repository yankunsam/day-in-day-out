package trie

type Trie struct {
	root  map[rune]*Trie
	isEnd bool
}

func Construtor() *Trie {
	return &Trie{map[rune]*Trie{}, false}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, s := range word {
		if _, ok := node.root[s]; !ok {
			node.root[s] = &Trie{map[rune]*Trie{}, false}
		}
		node = node.root[s]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, s := range word {
		if _, ok := node.root[s]; !ok {
			return false
		}
		node = node.root[s]
	}
	return node.isEnd
}
