package main

type History struct {
	history []Memento
}

func NewHistory() *History {
	return &History{make([]Memento, 0)}
}

func (h *History) Save(m Memento) {
	h.history = append(h.history, m)
}

func (h *History) Undo() Memento {
	if len(h.history) > 1 {
		n := len(h.history) - 1
		h.history = h.history[:n]
		return h.history[len(h.history)-1]
	} else {
		return Memento{}
	}
}
