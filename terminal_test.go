package main_test

import (
	cash "github.com/isaquesb/go-study-cash-terminal"
	"testing"
)

func TestTerminal_New(t *testing.T) {
	terminal := cash.NewTerminal()
	if terminal.Total() != 0 {
		t.Error("Total is not zero")
	}
}

func TestTerminal_Add(t *testing.T) {
	terminal := cash.NewTerminal()
	expected := 200
	if total := terminal.Add(10, 20); total != expected {
		t.Errorf("Total got %d, expected %d", total, expected)
	}
	if total := terminal.Add(5, 2); total != expected+10 {
		t.Errorf("Total got %d, expected %d", total, expected+10)
	}
}

func TestTerminal_Available(t *testing.T) {
	terminal := cash.NewTerminal()
	terminal.Add(5, 2)
	terminal.Add(100, 3)
	total := terminal.Add(10, 20)
	if total != 510 {
		t.Errorf("Total got %d, expected 510", total)
	}
	a := terminal.Available()
	for i, v := range []struct{ q, v int }{
		{3, 100},
		{20, 10},
		{2, 5},
	} {
		if a[i].Quantity != v.q {
			t.Errorf("Quantity got %d, expected %d", a[i].Quantity, v.q)
		}
		if a[i].Value != v.v {
			t.Errorf("Value got %d, expected %d", a[i].Value, v.q)
		}
	}
	s := terminal.GetSlot(10)
	if s.Quantity != 20 {
		t.Errorf("Quantity got %d, expected %d", s.Quantity, 20)
	}
	s = terminal.GetSlot(100)
	if s.Quantity != 3 {
		t.Errorf("Quantity got %d, expected %d", s.Quantity, 3)
	}
}

func TestTerminal_Overflow(t *testing.T) {
	terminal := cash.NewTerminal()
	terminal.Add(5, 2)
	terminal.Add(100, 3)
	terminal.Add(10, 20)
	sceneries, err := terminal.GetSceneriesFor(1000)
	if nil != sceneries {
		t.Error("Has scenery")
	}
	if "overflow required" != err.Error() {
		t.Error("Unexpected error: " + err.Error())
	}
}

func TestTerminal_NoSceneries(t *testing.T) {
	terminal := cash.NewTerminal()
	terminal.Add(5, 2)
	terminal.Add(100, 3)
	terminal.Add(10, 20)
	sceneries, err := terminal.GetSceneriesFor(11)
	if nil != sceneries {
		t.Error("Has scenery")
	}
	if "nothing sceneries" != err.Error() {
		t.Error("Unexpected error: " + err.Error())
	}
}

func TestTerminal_SimpleScenery(t *testing.T) {
	terminal := cash.NewTerminal()
	terminal.Add(5, 2)
	terminal.Add(100, 3)
	terminal.Add(10, 20)
	sceneries, err := terminal.GetSceneriesFor(320)
	if nil != err {
		t.Error(err.Error())
	}
	if l := len(sceneries); 1 != l {
		t.Errorf("got %d sceneries, expected 1", l)
		return
	}
	notes := sceneries[0].Cash
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

func TestTerminal_ManySceneries(t *testing.T) {
	terminal := cash.NewTerminal()
	terminal.Add(5, 1000)
	terminal.Add(100, 1000)
	terminal.Add(10, 1000)
	terminal.Add(50, 1000)
	terminal.Add(20, 1000)
	sceneries, err := terminal.GetSceneriesFor(845)
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
		notes := sceneries[j].Cash
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
