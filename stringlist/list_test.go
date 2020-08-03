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
	const listElt2 = "listelt2"
	eltAtIndex2 := theList.GetValue(2)
	if listElt2 != eltAtIndex2 {
		t.Error("Expected elt at index 2: ", listElt2, "but got ", eltAtIndex2)
	}

}
