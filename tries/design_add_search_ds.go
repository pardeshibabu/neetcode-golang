package tries

type WordDictionary struct {
	Values    [26]*WordDictionary
	EndOfWord bool
	root      *WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{
		Values:    [26]*WordDictionary{},
		EndOfWord: false,
		root:      &WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		ind := word[i] - 'a'
		if curr.Values[ind] == nil {
			curr.Values[ind] = &WordDictionary{}
		}
		curr = curr.Values[ind]
	}
	curr.EndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	// i := 0
	// var dfs func(int, *WordDictionary) bool
	// dfs = func(i int, root *WordDictionary) bool {
	// 	if word[i] == 46 {
	// 		for j := i + 1; j < len(word); j++ {

	// 			dfs(j,)
	// 		}
	// 	}
	// 	return true
	// }
	// for i := 0; i < len(word); i++ {
	// 	ind := word[i] - 'a'
	// 	if word[i] != 46 && curr.Values[ind] == nil {
	// 		return false
	// 	}
	// 	curr = curr.Values[ind]
	// }
	// curr.EndOfWord = true

	// return dfs(i, this.root)
	return true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
