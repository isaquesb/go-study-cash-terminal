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

func (request *Request) GetChoices(available Slots) ([]*Choice, error) {
	choices := make([]*Choice, 0)
	halfBase := 0
	halfCount := 0
	for i := 1; i <= SUGGEST_MAX; i++ {
		choice, rest := NewChoice(available, request.Value, halfBase)
		if 0 == rest {
			choices = append(choices, choice)
			halfBase = choice.Pieces[halfCount].Value
			halfCount++
		}
	}
	if 0 == len(choices) {
		return nil, errors.New("nothing choices")
	}
	return choices, nil
}
