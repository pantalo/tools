package nodisk

import (
	"golang.org/x/tools/pkg/lsp/foo"
)

func _() {
	foo.Foo() //@complete("F", Foo, IntFoo, StructFoo)
}
