package imports

import (
	"os"
	"testing"

	"golang.org/x/tools/pkg/testenv"
)

func TestMain(m *testing.M) {
	testenv.ExitIfSmallMachine()
	os.Exit(m.Run())
}
