package dllist

// Item это структура представляющая узел двусвязного списка.
type Item struct {
	value interface{}
	next  *Item
	prev  *Item
	list  *List
}

// Value возвращает значение элемента.
func (i *Item) Value() interface{} {
	return i.value
}

// Next возвращает следующий элемент списка.
func (i *Item) Next() *Item {
	return i.next
}

// Prev возвращает предыдущий элемент списка.
func (i *Item) Prev() *Item {
	return i.prev
}

// Remove удаляет элемент списка.
// Учитывает случаи когда в списке элемент крайний или последний.
func (i *Item) Remove() {
	if i.Prev() != nil {
		i.Prev().next = i.Next()
	}

	if i.Next() != nil {
		i.Next().prev = i.Prev()
	}

	if i.list.Len() == 1 {
		i.list.first = nil
		i.list.last = nil
	}

	i.prev = nil
	i.next = nil

	i.list.len--
}

// List это структура содержащая в себе элементы двусвязного списка.
type List struct {
	first *Item
	last  *Item
	len   int
}

// Len возвращает длину списка.
func (l *List) Len() int {
	return l.len
}

// First возвращает первый элемент списка.
func (l *List) First() *Item {
	return l.first
}

// Last возвращает последний элемент списка.
func (l *List) Last() *Item {
	return l.last
}

// PushFront добавляет новый элемент в начало списка.
// Учитывает случаи когда в списке один или несколько элементов.
func (l *List) PushFront(value interface{}) {
	fistItem := l.First()

	if fistItem == nil {
		l.first = &Item{value, nil, nil, l}
		l.last = l.first
	} else {
		newItem := &Item{value, fistItem, nil, l}
		l.first = newItem
	}

	l.len++
}

// PushFront добавляет новый элемент в конец списка.
// Учитывает случаи когда в списке один или несколько элементов.
func (l *List) PushBack(value interface{}) {
	lastItem := l.Last()

	if lastItem == nil {
		l.last = &Item{value, nil, nil, l}
		l.first = l.last
	} else {
		newItem := &Item{value, nil, lastItem, l}
		lastItem.next = newItem
		l.last = newItem
	}

	l.len++
}