package semantics

import (
	"github.com/sdkvictor/golang-compiler/directories"
)

func IdIsReserved(id string) bool {
	for _, f := range reservedFunctions {
		if id == f {
			return true
		}
	}
	return false
}

func checkVarDirReserved(vardir *directories.VarDirectory) (string, bool) {
	for _, f := range reservedFunctions {
		if vardir.Exists(f) {
			return f, false
		}
	}

	return "", true
}
