package app

import (
	"github.com/Tech-Arch1tect/OpenContainerForwarder/structs"
	"github.com/google/go-cmp/cmp"
)

// slicesEqual compares two slices of structs.ContainerExtracts
// cmp.Equal is used for the comparison however may be inappropriate for production use
func slicesEqual(a, b []structs.ContainerExtracts) bool {
	return cmp.Equal(a, b)
}
