package c

import "golang.org/x/tools/pkg/lsp/rename/b"

func _() {
	b.Hello() //@rename("Hello", "Goodbye")
}
