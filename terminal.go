package main

type Terminal struct {
	wallet *Wallet
}

func NewTerminal() *Terminal {
	return &Terminal{
		wallet: NewWallet(),
	}
}

func (t *Terminal) GetWallet() *Wallet {
	return t.wallet
}
