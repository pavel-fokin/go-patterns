package main

type List []string

type Tasks struct {
	list List
}

func NewTasks() *Tasks {
	return &Tasks{}
}

func (t *Tasks) Memento() Memento {
	return Memento{list: t.list}
}

func (t *Tasks) Restore(m Memento) {
	t.list = m.list
}

func (t *Tasks) Add(s string) {
	t.list = append(t.list, s)
}

func (t *Tasks) All() List {
	return t.list
}
