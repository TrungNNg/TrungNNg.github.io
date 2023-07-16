package main

import (
    "fmt"
)
var p = fmt.Println
func main() {
    trie := NewTrie()
    trie.Insert("abc")
    trie.Insert("ab")
    trie.Insert("def")
    trie.Insert("defgh")
   
    ans := trie.WithPrefix("")
    p(ans)
}

type Trie struct {
    Root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{Root: &TrieNode{m: map[rune]*TrieNode{}, isWord:false }}
}

type TrieNode struct {
    m map[rune]*TrieNode
    isWord bool
}

func (t *Trie) Insert(word string) {
    curr := t.Root
    for _, r := range word {
        if n, ok := curr.m[r]; ok {
            curr = n
        } else {
            curr.m[r] = &TrieNode{m: map[rune]*TrieNode{}, isWord:false}
            curr = curr.m[r]
        }
    }
    curr.isWord = true
}

func (t *Trie) WordExist(word string) bool {
    curr := t.Root
    for _, r := range word {
        if _, ok := curr.m[r]; !ok {
            return false
        }
        curr = curr.m[r]
    }
    return curr.isWord
}

// TODO
func (t *Trie) RemoveWord(word string) {
    
}

func (t *Trie) RemoveAllWithPrefix(prefix string) {
    
}

func (t *Trie) WithPrefix(prefix string) []string {
    curr := t.Root
    for _, r := range prefix {
        if _, ok := curr.m[r]; !ok {
            return []string{}
        }
        curr = curr.m[r]
    }
    ans := []string{}
    withPrefixDFS(curr, prefix, &ans)
    return ans
}

func withPrefixDFS (node *TrieNode, prefix string, ans *[]string) {
    if node.isWord {
        *ans = append(*ans, prefix)
    }
    for k, n := range node.m {
        withPrefixDFS(n, prefix + string(k), ans)
    }
}







