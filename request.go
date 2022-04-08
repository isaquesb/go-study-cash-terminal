package main

import (
	"errors"
)

const (
	SUGGEST_MAX = 3
)

type Request struct {
	Value int
}

func (request *Request) GetSceneries(available Slots) ([]*Scenery, error) {
	sceneries := make([]*Scenery, 0)
	halfBase := 0
	halfCount := 0
	for i := 1; i <= SUGGEST_MAX; i++ {
		scenery, rest := NewScenery(available, request.Value, halfBase)
		if 0 == rest {
			sceneries = append(sceneries, scenery)
			halfBase = scenery.Pieces[halfCount].Value
			halfCount++
		}
	}
	if 0 == len(sceneries) {
		return nil, errors.New("nothing sceneries")
	}
	return sceneries, nil
}
