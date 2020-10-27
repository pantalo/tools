package errors

import (
	"golang.org/x/tools/pkg/lsp/types"
)

func _() {
	bob.Bob() //@complete(".")
	types.b //@complete(" //", Bob_interface)
}
