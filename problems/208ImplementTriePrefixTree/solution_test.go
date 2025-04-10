package trieprefixtree

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tests := []struct {
		name       string
		operations []struct {
			op     string
			word   string
			expect bool
		}
	}{
		{
			name: "basic operations",
			operations: []struct {
				op     string
				word   string
				expect bool
			}{
				{op: "insert", word: "apple"},
				{op: "search", word: "apple", expect: true},
				{op: "search", word: "app", expect: false},
				{op: "startsWith", word: "app", expect: true},
				{op: "insert", word: "app"},
				{op: "search", word: "app", expect: true},
			},
		},
		{
			name: "empty trie",
			operations: []struct {
				op     string
				word   string
				expect bool
			}{
				{op: "search", word: "apple", expect: false},
				{op: "startsWith", word: "app", expect: false},
			},
		},
		{
			name: "multiple words",
			operations: []struct {
				op     string
				word   string
				expect bool
			}{
				{op: "insert", word: "apple"},
				{op: "insert", word: "banana"},
				{op: "insert", word: "orange"},
				{op: "search", word: "apple", expect: true},
				{op: "search", word: "banana", expect: true},
				{op: "search", word: "orange", expect: true},
				{op: "search", word: "app", expect: false},
				{op: "startsWith", word: "app", expect: true},
				{op: "startsWith", word: "ban", expect: true},
				{op: "startsWith", word: "ora", expect: true},
				{op: "startsWith", word: "xyz", expect: false},
			},
		},
		{
			name: "prefix of existing word",
			operations: []struct {
				op     string
				word   string
				expect bool
			}{
				{op: "insert", word: "apple"},
				{op: "search", word: "app", expect: false},
				{op: "startsWith", word: "app", expect: true},
				{op: "insert", word: "app"},
				{op: "search", word: "app", expect: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for _, op := range tt.operations {
				switch op.op {
				case "insert":
					trie.Insert(op.word)
				case "search":
					result := trie.Search(op.word)
					if result != op.expect {
						t.Errorf("Search(%q) = %v, want %v", op.word, result, op.expect)
					}
				case "startsWith":
					result := trie.StartsWith(op.word)
					if result != op.expect {
						t.Errorf("StartsWith(%q) = %v, want %v", op.word, result, op.expect)
					}
				}
			}
		})
	}
}
