package trie

import "testing"

func TestSearch(t *testing.T) {
	tr := Construtor()
	s1 := "hello"
	s2 := "world"
	tr.Insert(s1)
	tr.Insert(s2)
	if !tr.Search("hello") || !tr.Search(s2) {
		t.Errorf("both %s and %s should exist!", s1, s2)
	}
	if tr.Search("he") {
		t.Errorf("he shoud not exist!")
	}

}
