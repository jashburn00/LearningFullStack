package main

//my dope custom list
//here is a generic "list" type, which can hold any data type
type DopeList[T any] struct {
	items []T
}

//method to add an item to the list
func (l *DopeList[T]) Add(item T) {
	l.items = append(l.items, item)
}
