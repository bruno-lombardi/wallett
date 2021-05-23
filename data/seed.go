package data

import (
	"fmt"
	"math/rand"
	"time"
	"wallett/generators"
	"wallett/models"
)

type Wallet = models.Wallet
type User = models.User
type Transaction = models.Transaction

func SeedInitialWallets(wallets *[]Wallet, users *[]User) {
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
	for i, wallet := range *wallets {
		rand.Seed(time.Now().UnixNano())
		amounts := generators.FloatsInRange(-12000.0, 13000.0, 200)

		for _, amount := range amounts {
			wallet.AddTransaction(Transaction{
				ID:           generators.ID("trx"),
				WalletID:     wallet.ID,
				Amount:       amount,
				CurrencyCode: "BRL",
			})
			(*wallets)[i] = wallet
		}
	}
}
