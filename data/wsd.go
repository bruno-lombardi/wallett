package data

import (
	"fmt"
	"wallett/persistence"
)

type Persistable interface {
	Persist() error
}

type WSD struct {
	Persistable
	Wallets *[]Wallet
	Users   *[]User
}

func GetWSD() *WSD {
	var wsd *WSD = &WSD{}
	var wallets *[]Wallet = &[]Wallet{}
	var users *[]User = &[]User{}

	wsd.Users = users
	wsd.Wallets = wallets

	if err := persistence.ReadAndDecodeFile("wsd.dat", &wsd); err != nil {
		fmt.Println("--- Wallett ---")
		fmt.Println("No wallett data found... initializing wallett")
		fmt.Println(err)
	}

	// Seed wallets
	if len(*wallets) == 0 {
		SeedInitialWallets(wallets, users)
	}
	return wsd
}

func (wsd *WSD) PersistWSD() {
	if err := persistence.WriteAndEncodeFile("wsd.dat", &wsd); err != nil {
		fmt.Println("We could not save wallets...")
		fmt.Println(err)
	} else {
		fmt.Println("Success saving wallett data!")
	}
}
