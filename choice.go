package main

import "sort"

type Choice struct {
	Pieces Slots
}

func NewChoice(available Slots, value, halfBase int) (*Choice, int) {
	sortAvailable(available)
	choice := &Choice{
		Pieces: make(Slots, 0),
	}
	rest := value
	for _, availableSlot := range available {
		piece := availableSlot.getPiece(rest, halfBase)
		if nil == piece {
			continue
		}
		rest -= piece.Value * piece.Quantity
		choice.Pieces = append(choice.Pieces, piece)
	}
	return choice, rest
}

func sortAvailable(available Slots) {
	sort.SliceStable(available, func(i, j int) bool {
		return available[i].Value > available[j].Value
	})
}
