package stringlist_test

import (
	"testing"

	"github.com/osscda/lets-go-project/stringlist"
)

func TestNewList(t *testing.T) {
	const listElt0 = "listelt0"
	theList := stringlist.NewList(listElt0)
	eltAtIndex0 := theList.GetValue(0)
	if listElt0 != eltAtIndex0 {
		t.Error("Expected elt at index 0: ", listElt0, "but got ", eltAtIndex0)
	}
	const listElt1 = "listelt1"
	theList.Append(listElt1)
	eltAtIndex1 := theList.GetValue(1)
	if listElt1 != eltAtIndex1 {
		t.Error("Expected elt at index 1: ", listElt1, "but got ", eltAtIndex1)
	}
}

func TestReplace(t *testing.T) {
	elements := []string{
		"listelt1",
		"listelt2",
		"listelt3",
	}
	theList := stringlist.NewList("listelt0")
	// this loops through the entire elements list ("slice" in Go)
	//
	// add all of them to our list
	for _, elt := range elements {
		theList.Append(elt)
	}

	// now replace the second item
	theList.Replace(1, "newlistelt1")

	newSecondElt := theList.GetValue(1)
	if newSecondElt != "newlistelt1" {
		t.Errorf(
			"Expected 'newlistelt1' in index 1, got %s",
			newSecondElt,
		)
	}

}
