package main

import (
	"errors"
	"fmt"
	"sort"
)

type Scenery struct {
	Cash []*Node
}

type Node struct {
	Value    int
	Quantity int
}

type Terminal struct {
	slots     []*Node
	slotIndex map[int]*Node
	total     int
}

func NewTerminal() *Terminal {
	return &Terminal{
		slots:     make([]*Node, 0),
		slotIndex: map[int]*Node{},
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

func (t *Terminal) Available() []*Node {
	sort.SliceStable(t.slots, func(i, j int) bool {
		return t.slots[i].Value > t.slots[j].Value
	})
	return t.slots
}

func (t *Terminal) GetSlot(value int) *Node {
	if nil == t.slotIndex[value] {
		n := &Node{
			Value:    value,
			Quantity: 0,
		}
		t.slotIndex[value] = n
		t.slots = append(t.slots, n)
	}
	return t.slotIndex[value]
}

func (t *Terminal) getSceneryFor(required, half int) (*Scenery, error) {
	newScenery := &Scenery{
		Cash: make([]*Node, 0),
	}
	left := required
	fmt.Println("============================")
	fmt.Println(half)
	for _, scenery := range t.Available() {
		if left < scenery.Value {
			continue
		}
		notes := left / scenery.Value
		if half > 0 && scenery.Value >= half {
			notes = notes / 2
		}
		if notes > scenery.Quantity {
			notes = scenery.Quantity
		}
		node := &Node{
			Value:    scenery.Value,
			Quantity: notes,
		}
		left -= scenery.Value * notes
		newScenery.Cash = append(newScenery.Cash, node)
	}
	if left == 0 {
		return newScenery, nil
	}
	min := newScenery.Cash[len(newScenery.Cash)-1]
	slot := t.GetSlot(min.Value)
	return nil, errors.New(fmt.Sprintf(
		"need %d, used %d x R$ %d for total %d notes",
		left,
		min.Quantity,
		min.Value,
		slot.Quantity,
	))
}

func (t *Terminal) GetSceneriesFor(required int) ([]*Scenery, error) {
	if required > t.Total() {
		return nil, errors.New("overflow required")
	}
	sceneries := make([]*Scenery, 0)
	half := 0
	halfCount := 0
	for i := 1; i <= 3; i++ {
		newScenery, err := t.getSceneryFor(required, half)
		if nil == err {
			sceneries = append(sceneries, newScenery)
			half = newScenery.Cash[halfCount].Value
			halfCount++
		} else {
			fmt.Println(err.Error())
		}
	}
	if 0 == len(sceneries) {
		return nil, errors.New("nothing sceneries")
	}
	return sceneries, nil
}
