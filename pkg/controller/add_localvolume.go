package controller

import (
	"github.com/HwameiStor/local-storage/pkg/controller/localvolume"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, localvolume.Add)
}