package tree_sitter_strace_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-strace"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_strace.Language())
	if language == nil {
		t.Errorf("Error loading Strace grammar")
	}
}
