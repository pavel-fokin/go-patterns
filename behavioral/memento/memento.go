package main

type Memento struct {
	list List
}

func (m *Memento) List() List {
	return m.list
}
