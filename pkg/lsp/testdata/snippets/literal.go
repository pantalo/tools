package snippets

import (
	"golang.org/x/tools/pkg/lsp/signature"
	t "golang.org/x/tools/pkg/lsp/types"
)

type structy struct {
	x signature.MyType
}

func X(_ map[signature.Alias]t.CoolAlias) (map[signature.Alias]t.CoolAlias) {
	return nil
}

func _() {
	X() //@signature(")", "X(_ map[signature.Alias]t.CoolAlias) map[signature.Alias]t.CoolAlias", 0)
	_ = signature.MyType{} //@item(literalMyType, "signature.MyType{}", "", "var")
	s := structy{
		x: //@snippet(" //", literalMyType, "signature.MyType{\\}", "signature.MyType{\\}")
	}
}