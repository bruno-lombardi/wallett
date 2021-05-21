package main

import (
	"fmt"
	"math/rand"
	"time"
	"wallett/generators"
	"wallett/models"
	"wallett/persistence"
)

type Wallet = models.Wallet
type User = models.User
type Transaction = models.Transaction

type WSD struct {
	Wallets *[]Wallet
	Users   *[]User
}

func main() {
	var wsd WSD = WSD{}
	var wallets *[]Wallet = &[]Wallet{}
	var users *[]User = &[]User{}

	wsd.Users = users
	wsd.Wallets = wallets

	if err := persistence.ReadAndDecodeFile("wsd.dat", &wsd); err != nil {
		fmt.Println("--- Wallett ---")
		fmt.Println("No wallett data found... initializing wallett")
		fmt.Println(err)
	}

	if len(*wallets) == 0 {
		for i := 0; i < 100; i++ {
			user := User{
				ID:       generators.ID("u"),
				Email:    fmt.Sprintf("%v@%v.com", generators.RandomString(30), generators.RandomString(10)),
				Name:     fmt.Sprintf("Account #%d", i),
				Password: generators.RandomString(12),
			}
			wallet := Wallet{
				ID:           generators.ID("wa"),
				UserID:       user.ID,
				CurrencyCode: "BRL",
				Balance:      0,
				Transactions: []Transaction{},
			}
			*wallets = append(*wallets, wallet)
			*users = append(*users, user)
		}
	}

	for i, wallet := range *wsd.Wallets {
		rand.Seed(time.Now().UnixNano())
		amounts := generators.FloatsInRange(-12000.0, 13000.0, 200)

		for _, amount := range amounts {
			wallet.AddTransaction(Transaction{
				ID:           generators.ID("trx"),
				WalletID:     wallet.ID,
				Amount:       amount,
				CurrencyCode: "BRL",
			})
			(*wsd.Wallets)[i] = wallet
		}
	}

	fmt.Println("Saving wallets...")
	if err := persistence.WriteAndEncodeFile("wsd.dat", &wsd); err != nil {
		fmt.Println("We could not save wallets...")
		fmt.Println(err)
	} else {
		fmt.Println("Success saving wallett data!")
		fmt.Printf("> Wallets Count: %v\n", len(*wsd.Wallets))
		var walletHoldings float64
		for _, wallet := range *wsd.Wallets {
			fmt.Printf("===================\n")
			fmt.Printf("Wallet ID: %v\n", wallet.ID)
			fmt.Printf("Balance: %.2f\n", wallet.Balance)
			fmt.Printf("Transactions Count: %v\n", len(wallet.Transactions))
			walletHoldings += wallet.Balance
		}
		fmt.Printf("===================\n")
		fmt.Printf("Total Wallett Holdings: %.2f BRL\n", walletHoldings)
	}

}
