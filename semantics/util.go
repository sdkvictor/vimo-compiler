package sem

import (
	"github.com/sdkvictor/golang-compiler/directories"
)

func idIsReserved(id string) bool {
	return id == "pow" || id == "sqrt" || id == "loadImage" || id == "keyPressed" || id == "checkCollision"
}

func checkVarDirReserved(vardir *directories.VarDirectory) (string, bool) {
	for _, f := range reservedFunctions {
		if vardir.Exists(f) {
			return f, false
		}
	}

	return "", true
}
