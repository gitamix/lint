package testdata

import (
	"path/filepath"
	"runtime"
)

// PathToFile wraps the provided filename
// with path to current 'testdata/' directory and returns it.
//
// Panics on failed to get the runtime caller.
//
// The method does not check whether the file exists.
func PathToFile(fname string) string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("failed to resolve runtime caller")
	}
	return filepath.Dir(file) + "/testdata/" + fname
}
