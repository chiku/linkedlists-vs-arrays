package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_array_list_is_empty_on_creation(t *testing.T) {
	assert.True(t, NewArraylist(3).IsEmpty(), "New list must be empty")
}

func Test_array_list_has_reports_items_as_zero_before_insertion(t *testing.T) {
	list := NewArraylist(3)

	assert.Equal(t, 0, list.Get(0))
	assert.Equal(t, 0, list.Get(1))
}

func Test_array_list_has_items_when_first_item_is_inserted(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)

	assert.False(t, list.IsEmpty(), "List with items must not be empty")
	assert.Equal(t, 10, list.Get(0))
}

func Test_array_list_has_items_when_second_item_is_inserted_at_end(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 1)

	assert.False(t, list.IsEmpty(), "List with items must not be empty")
	assert.Equal(t, 10, list.Get(0))
	assert.Equal(t, 20, list.Get(1))
}

func Test_array_list_has_items_when_item_is_inserted_beyond_end(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 2)

	assert.False(t, list.IsEmpty(), "List with items must not be empty")
	assert.Equal(t, 10, list.Get(0))
	assert.Equal(t, 20, list.Get(1))
}

func Test_array_list_has_items_when_second_item_is_inserted_at_start(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 0)

	assert.False(t, list.IsEmpty(), "List with items must not be empty")
	assert.Equal(t, 20, list.Get(0))
	assert.Equal(t, 10, list.Get(1))
}

func Test_array_list_has_items_when_third_item_is_inserted_in_middle(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 1)
	list.Insert(30, 1)

	assert.False(t, list.IsEmpty(), "List with items must not be empty")
	assert.Equal(t, 10, list.Get(0))
	assert.Equal(t, 30, list.Get(1))
	assert.Equal(t, 20, list.Get(2))
}

func Test_array_list_has_a_string_representation(t *testing.T) {
	list := NewArraylist(3)
	list.Insert(10, 0)
	list.Insert(20, 1)
	list.Insert(30, 1)

	assert.Equal(t, "[ <10> -> <30> -> <20> -> NULL ]", list.String())
}
