package main

import (
	"errors"
	"fmt"
	"sort"
)

const (
	SUGGEST_MAX = 3
)

type Scenery struct {
	Pieces []*Slot
}

type Slot struct {
	Value    int
	Quantity int
}

type Terminal struct {
	slots     []*Slot
	slotIndex map[int]*Slot
	total     int
}

func NewTerminal() *Terminal {
	return &Terminal{
		slots:     make([]*Slot, 0),
		slotIndex: map[int]*Slot{},
		total:     0,
	}
}

func (t *Terminal) Total() int {
	return t.total
}

func (t *Terminal) Add(value, quantity int) int {
	slot := t.GetSlot(value)
	slot.Quantity += quantity
	t.total += quantity * value
	return t.total
}

func (t *Terminal) Available() []*Slot {
	sort.SliceStable(t.slots, func(i, j int) bool {
		return t.slots[i].Value > t.slots[j].Value
	})
	return t.slots
}

func (t *Terminal) GetSlot(value int) *Slot {
	if nil == t.slotIndex[value] {
		n := &Slot{
			Value:    value,
			Quantity: 0,
		}
		t.slotIndex[value] = n
		t.slots = append(t.slots, n)
	}
	return t.slotIndex[value]
}

func (t *Terminal) GetSceneriesFor(required int) ([]*Scenery, error) {
	if required > t.Total() {
		return nil, errors.New("overflow required")
	}
	sceneries := make([]*Scenery, 0)
	halfBase := 0
	halfCount := 0
	for i := 1; i <= SUGGEST_MAX; i++ {
		newScenery, err := t.getSceneryFor(required, halfBase)
		if nil == err {
			sceneries = append(sceneries, newScenery)
			halfBase = newScenery.Pieces[halfCount].Value
			halfCount++
		}
	}
	if 0 == len(sceneries) {
		return nil, errors.New("nothing sceneries")
	}
	return sceneries, nil
}

func (t *Terminal) getSceneryFor(required, halfBase int) (*Scenery, error) {
	scenery := &Scenery{
		Pieces: make([]*Slot, 0),
	}
	need := required
	for _, terminalSlot := range t.Available() {
		slot := t.getSlotPiece(terminalSlot, need, halfBase)
		if slot == nil {
			continue
		}
		need -= slot.Value * slot.Quantity
		scenery.Pieces = append(scenery.Pieces, slot)
	}
	if need == 0 {
		return scenery, nil
	}
	min := scenery.Pieces[len(scenery.Pieces)-1]
	slot := t.GetSlot(min.Value)
	return nil, errors.New(fmt.Sprintf(
		"need %d, used %d x R$ %d for total %d notes",
		need,
		min.Quantity,
		min.Value,
		slot.Quantity,
	))
}

func (t *Terminal) getSlotPiece(scenery *Slot, need, halfBase int) *Slot {
	if need < scenery.Value {
		return nil
	}
	notes := need / scenery.Value
	if halfBase > 0 && scenery.Value >= halfBase {
		notes = notes / 2
	}
	if notes > scenery.Quantity {
		notes = scenery.Quantity
	}
	return &Slot{
		Value:    scenery.Value,
		Quantity: notes,
	}
}
