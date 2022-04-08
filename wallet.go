package main

import "errors"

type Wallet struct {
	slots Slots
	index map[int]*Slot
	total int
}

func NewWallet() *Wallet {
	return &Wallet{
		index: map[int]*Slot{},
		total: 0,
	}
}

func (w *Wallet) Available() Slots {
	return w.slots
}

func (w *Wallet) Add(value, quantity int) int {
	slot := w.getSlot(value)
	slot.Quantity += quantity
	w.total += quantity * value
	return w.total
}

func (w *Wallet) Total() int {
	return w.total
}

func (w *Wallet) Request(value int) (*Request, error) {
	if value > w.Total() {
		return nil, errors.New("request overflowed")
	}
	return &Request{value}, nil
}

func (w *Wallet) getSlot(value int) *Slot {
	if nil == w.index[value] {
		slot := &Slot{
			Value:    value,
			Quantity: 0,
		}
		w.index[value] = slot
		w.slots = append(w.slots, slot)
	}
	return w.index[value]
}
