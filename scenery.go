package main

import "sort"

type Scenery struct {
	Pieces Slots
}

func NewScenery(available Slots, value, halfBase int) (*Scenery, int) {
	sortAvailable(available)
	scenery := &Scenery{
		Pieces: make(Slots, 0),
	}
	rest := value
	for _, availableSlot := range available {
		piece := availableSlot.getPiece(rest, halfBase)
		if nil == piece {
			continue
		}
		rest -= piece.Value * piece.Quantity
		scenery.Pieces = append(scenery.Pieces, piece)
	}
	return scenery, rest
}

func sortAvailable(available Slots) {
	sort.SliceStable(available, func(i, j int) bool {
		return available[i].Value > available[j].Value
	})
}
