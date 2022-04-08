package main_test

import (
	cash "github.com/isaquesb/go-study-cash-terminal"
	"testing"
)

func TestTerminal_New(t *testing.T) {
	terminal := cash.NewTerminal()
	if terminal.GetWallet().Total() != 0 {
		t.Error("Wallet not empty")
	}
}
