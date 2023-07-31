package tries

type Trie struct {
	Values [26]*Trie
	Flag   bool
	root   *Trie
}

// func Constructor() Trie {
// 	return Trie{
// 		Values: [26]*Trie{},
// 		Flag:   false,
// 		root:   &Trie{},
// 	}
// }

func (this *Trie) Insert(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		ind := word[i] - 'a'
		if cur.Values[ind] == nil {
			cur.Values[ind] = &Trie{}
		}
		cur = cur.Values[ind]
	}
	cur.Flag = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for i := 0; i < len(word); i++ {
		ind := word[i] - 'a'
		if cur.Values[ind] == nil {
			return false
		}
		cur = cur.Values[ind]
	}
	return cur.Flag
}

func (this *Trie) StartsWith(word string) bool {
	cur := this.root
	for i := 0; i < len(word); i++ {
		ind := word[i] - 'a'
		if cur.Values[ind] == nil {
			return false
		}
		cur = cur.Values[ind]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
