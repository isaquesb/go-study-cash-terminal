package main_test

import (
	cash "github.com/isaquesb/go-study-cash-terminal"
	"testing"
)

func TestWallet_New(t *testing.T) {
	wallet := cash.NewWallet()
	if wallet.Total() != 0 {
		t.Error("Total is not zero")
	}
}

func TestWallet_Add(t *testing.T) {
	wallet := cash.NewWallet()
	expected := 200
	if total := wallet.Add(10, 20); total != expected {
		t.Errorf("Total got %d, expected %d", total, expected)
	}
	if total := wallet.Add(5, 2); total != expected+10 {
		t.Errorf("Total got %d, expected %d", total, expected+10)
	}
}

func TestWallet_Available(t *testing.T) {
	wallet := cash.NewWallet()
	wallet.Add(5, 2)
	wallet.Add(100, 3)
	total := wallet.Add(10, 20)
	if total != 510 {
		t.Errorf("Total got %d, expected 510", total)
	}
	a := wallet.Available()
	for i, v := range []struct{ q, v int }{
		{2, 5},
		{3, 100},
		{20, 10},
	} {
		if a[i].Quantity != v.q {
			t.Errorf("Quantity got %d, expected %d", a[i].Quantity, v.q)
		}
		if a[i].Value != v.v {
			t.Errorf("Value got %d, expected %d", a[i].Value, v.q)
		}
	}
}
