package main

type Slots = []*Slot

type Slot struct {
	Value    int
	Quantity int
}

func (s *Slot) getPiece(need, halfBase int) *Slot {
	if need < s.Value {
		return nil
	}
	notes := need / s.Value
	if halfBase > 0 && s.Value >= halfBase {
		notes = notes / 2
	}
	if notes > s.Quantity {
		notes = s.Quantity
	}
	return &Slot{
		Value:    s.Value,
		Quantity: notes,
	}
}
