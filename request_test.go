package main_test

import (
	cash "github.com/isaquesb/go-study-cash-terminal"
	"testing"
)

func TestRequest_Overflow(t *testing.T) {
	wallet := cash.NewWallet()
	wallet.Add(5, 2)
	wallet.Add(100, 3)
	wallet.Add(10, 20)
	request, err := wallet.Request(1000)
	if nil != request {
		t.Error("Has request")
	}
	if "request overflowed" != err.Error() {
		t.Error("Unexpected error: " + err.Error())
		return
	}
}

func TestRequest_WithoutChoices(t *testing.T) {
	wallet := cash.NewWallet()
	wallet.Add(5, 2)
	wallet.Add(100, 3)
	wallet.Add(10, 20)
	request, err := wallet.Request(11)
	choices, err := request.GetChoices(wallet.Available())
	if nil != choices {
		t.Error("Has choice")
	}
	if "nothing choices" != err.Error() {
		t.Error("Unexpected error: " + err.Error())
	}
}

func TestRequest_SimpleChoice(t *testing.T) {
	wallet := cash.NewWallet()
	wallet.Add(5, 2)
	wallet.Add(100, 3)
	wallet.Add(10, 20)
	request, err := wallet.Request(320)
	choices, err := request.GetChoices(wallet.Available())
	if nil != err {
		t.Error(err.Error())
	}
	if l := len(choices); 1 != l {
		t.Errorf("got %d choices, expected 1", l)
		return
	}
	notes := choices[0].Pieces
	if num := len(notes); num != 2 {
		t.Errorf("got %d notes, expected 2", num)
	}
	for i, v := range []struct{ q, v int }{
		{3, 100},
		{2, 10},
	} {
		if notes[i].Quantity != v.q {
			t.Errorf("Quantity got %d, expected %d", notes[i].Quantity, v.q)
		}
		if notes[i].Value != v.v {
			t.Errorf("Value got %d, expected %d", notes[i].Value, v.q)
		}
	}
}

func TestRequest_ManyChoices(t *testing.T) {
	wallet := cash.NewWallet()
	wallet.Add(5, 1000)
	wallet.Add(100, 1000)
	wallet.Add(10, 1000)
	wallet.Add(50, 1000)
	wallet.Add(20, 1000)
	request, _ := wallet.Request(845)
	choices, err := request.GetChoices(wallet.Available())
	if nil != err {
		t.Error(err.Error())
	}
	if l := len(choices); 3 != l {
		t.Errorf("got %d choices, expected 3", l)
		return
	}
	cases := []struct{ n []struct{ q, v int } }{
		{
			[]struct{ q, v int }{
				{8, 100},
				{2, 20},
				{1, 5},
			},
		},
		{
			[]struct{ q, v int }{
				{4, 100},
				{8, 50},
				{2, 20},
				{1, 5},
			},
		},
		{
			[]struct{ q, v int }{
				{4, 100},
				{4, 50},
				{12, 20},
				{1, 5},
			},
		},
	}
	for j, currentCase := range cases {
		notes := choices[j].Pieces
		for i, v := range currentCase.n {
			if notes[i].Quantity != v.q {
				t.Errorf("Choice %d: Quantity got %d x %d, expected %d", j, notes[i].Quantity, notes[i].Value, v.q)
			}
			if notes[i].Value != v.v {
				t.Errorf("Choice %d: Value got %d, expected %d", j, notes[i].Value, v.q)
			}
		}
	}
}
