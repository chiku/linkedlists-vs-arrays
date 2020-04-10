package arraylist

import (
	"testing"
)

func TestArrayListEmptyWhenCreated(t *testing.T) {
	if !NewArraylist(3).IsEmpty() {
		t.Errorf("New list must be empty")
	}
}

func TestArrayListHasReportsItemsAsZeroBeforeInsertion(t *testing.T) {
	list := NewArraylist(3)

	assertEqualInt(t, list.Get(0), 0)
	assertEqualInt(t, list.Get(1), 0)
}

func TestArrayListHasItemsWhenFirstItemInserted(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)

	if list.IsEmpty() {
		t.Errorf("List with items must not be empty")
	}
	assertEqualInt(t, list.Get(0), 10)
}

func TestArrayListHasItemsWhenSecondItemInsertedAtEnd(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 1)

	if list.IsEmpty() {
		t.Errorf("List with items must not be empty")
	}
	assertEqualInt(t, list.Get(0), 10)
	assertEqualInt(t, list.Get(1), 20)
}

func TestArrayListHasItemsWhenItemInsertedBeyondEnd(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 2)

	if list.IsEmpty() {
		t.Errorf("List with items must not be empty")
	}
	assertEqualInt(t, list.Get(0), 10)
	assertEqualInt(t, list.Get(1), 20)
}

func TestArrayListHasItemsWhenSecondItemIsInsertedAtStart(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 0)

	if list.IsEmpty() {
		t.Errorf("List with items must not be empty")
	}
	assertEqualInt(t, list.Get(0), 20)
	assertEqualInt(t, list.Get(1), 10)
}

func TestArrayListHasItemsWhenThirdItemInsertedInMiddle(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 1)
	list.Insert(30, 1)

	if list.IsEmpty() {
		t.Errorf("List with items must not be empty")
	}
	assertEqualInt(t, list.Get(0), 10)
	assertEqualInt(t, list.Get(1), 30)
	assertEqualInt(t, list.Get(2), 20)
}

func TestArrayListHasStringRepresentation(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 1)
	list.Insert(30, 1)

	assertEqualString(t, list.String(), "[ <10> -> <30> -> <20> -> NULL ]")
}

func assertEqualInt(t *testing.T, actual, expected int) {
	t.Helper()
	if expected != actual {
		t.Errorf("%d != %d", actual, expected)
	}
}

func assertEqualString(t *testing.T, actual, expected string) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s != %s", actual, expected)
	}
}
