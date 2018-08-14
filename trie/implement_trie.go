type Trie struct {
    children map[string]*Trie
	endOfWord bool 
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie {
        children: make(map[string]*Trie, 0),
        endOfWord: false,
    }
}

func addNode(this *Trie, value string) *Trie {
	return &Trie {
        children: make(map[string]*Trie, 0),
        endOfWord: false,
    }
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    current := this 
	for _, v := range word {
		if _, ok  := current.children[string(v)]; !ok {
			//node not in children - add one.
			current.children[string(v)] = addNode(this, string(v))
            current = current.children[string(v)]
            fmt.Printf("current: %+v\n", current)
		} else {
			//already child - move to next
		    current = current.children[string(v)]
	    }
	}
	current.endOfWord = true 
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    current := this 
	for _, v := range word {
		if _, ok  := current.children[string(v)]; !ok {
			//node not in children - add one.
            fmt.Printf("NOT FOUND: %+v\n", string(v))
            return false 
		} else {
			//already child - move to next
		    current = current.children[string(v)]
	    }
	}
	return current.endOfWord 		
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    current := this 
	for _, v := range prefix {
		if _, ok  := current.children[string(v)]; !ok {
			//node not in children - add one.
            fmt.Printf("NOT FOUND: %+v\n", string(v))
            return false 
		} else {
			//already child - move to next
		    current = current.children[string(v)]
	    }
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


