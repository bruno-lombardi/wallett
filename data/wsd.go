package data

import (
	"fmt"
	"wallett/infra/persistence"
)

type WSD struct {
	Wallets         *[]Wallet
	Users           *[]User
	storageFilePath string
}

func NewWSD(storageFilePath string) *WSD {
	var wsd *WSD = &WSD{
		storageFilePath: storageFilePath,
	}
	var wallets *[]Wallet = &[]Wallet{}
	var users *[]User = &[]User{}

	wsd.Users = users
	wsd.Wallets = wallets

	if err := persistence.ReadAndDecodeFile(storageFilePath, &wsd); err != nil {
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

func (wsd *WSD) PersistWSD() (err error) {
	if err := persistence.WriteAndEncodeFile(wsd.storageFilePath, &wsd); err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

func (wsd *WSD) ClearWSD() (err error) {
	var wallets *[]Wallet = &[]Wallet{}
	var users *[]User = &[]User{}

	wsd.Users = users
	wsd.Wallets = wallets

	err = persistence.DeleteFile(wsd.storageFilePath)
	return err
}
