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
	quantity := need / s.Value
	if halfBase > 0 && s.Value >= halfBase {
		quantity = quantity / 2
	}
	if quantity > s.Quantity {
		quantity = s.Quantity
	}
	return &Slot{
		Value:    s.Value,
		Quantity: quantity,
	}
}
