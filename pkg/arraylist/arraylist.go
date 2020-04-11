package arraylist

import "strconv"

type Arraylist struct {
	capacity int
	tail     int
	values   []int
}

func NewArraylist(capacity int) Arraylist {
	return Arraylist{
		capacity: capacity,
		tail:     -1,
		values:   make([]int, capacity, capacity),
	}
}

func (list Arraylist) IsEmpty() bool {
	return list.tail == -1
}

func (list Arraylist) Get(position int) int {
	return list.values[position]
}

func (list *Arraylist) Insert(value int, position int) {
	list.tail++

	if position > list.tail {
		position = list.tail
	}

	list.values = append(list.values, 0)
	copy(list.values[position+1:], list.values[position:])
	list.values[position] = value
}

func (list Arraylist) String() string {
	result := "[ "

	for i := 0; i < list.capacity; i++ {
		result += "<"
		result += strconv.Itoa(list.values[i])
		result += "> -> "
	}

	result += "NULL ]"

	return result
}
