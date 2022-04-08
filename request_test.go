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

func TestRequest_WithoutSceneries(t *testing.T) {
	wallet := cash.NewWallet()
	wallet.Add(5, 2)
	wallet.Add(100, 3)
	wallet.Add(10, 20)
	request, err := wallet.Request(11)
	sceneries, err := request.GetSceneries(wallet.Available())
	if nil != sceneries {
		t.Error("Has scenery")
	}
	if "nothing sceneries" != err.Error() {
		t.Error("Unexpected error: " + err.Error())
	}
}

func TestRequest_SimpleScenery(t *testing.T) {
	wallet := cash.NewWallet()
	wallet.Add(5, 2)
	wallet.Add(100, 3)
	wallet.Add(10, 20)
	request, err := wallet.Request(320)
	sceneries, err := request.GetSceneries(wallet.Available())
	if nil != err {
		t.Error(err.Error())
	}
	if l := len(sceneries); 1 != l {
		t.Errorf("got %d sceneries, expected 1", l)
		return
	}
	notes := sceneries[0].Pieces
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

func TestRequest_ManySceneries(t *testing.T) {
	wallet := cash.NewWallet()
	wallet.Add(5, 1000)
	wallet.Add(100, 1000)
	wallet.Add(10, 1000)
	wallet.Add(50, 1000)
	wallet.Add(20, 1000)
	request, _ := wallet.Request(845)
	sceneries, err := request.GetSceneries(wallet.Available())
	if nil != err {
		t.Error(err.Error())
	}
	if l := len(sceneries); 3 != l {
		t.Errorf("got %d sceneries, expected 3", l)
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
		notes := sceneries[j].Pieces
		for i, v := range currentCase.n {
			if notes[i].Quantity != v.q {
				t.Errorf("Scenery %d: Quantity got %d x %d, expected %d", j, notes[i].Quantity, notes[i].Value, v.q)
			}
			if notes[i].Value != v.v {
				t.Errorf("Scenery %d: Value got %d, expected %d", j, notes[i].Value, v.q)
			}
		}
	}
}
