package sqlite

import (
	"fmt"
	"math"
	"wallett/domain/models"
	"wallett/infra/generators"

	"gorm.io/gorm"
)

type SQLiteWalletRepository struct {
	db *gorm.DB
}

type SQLiteTransaction struct {
	gorm.Model
	ID              string  `gorm:"primaryKey;type:VARCHAR(12);not null;unique"`
	WalletID        string  `gorm:"type:VARCHAR(12)"`
	Amount          float32 `gorm:"type:DECIMAL(10, 2)"`
	PreviousBalance float32 `gorm:"type:DECIMAL(10, 2)"`
	CurrencyCode    string  `gorm:"type:VARCHAR(3)"`
}

type SQLiteWallet struct {
	gorm.Model
	ID           string `gorm:"primaryKey;type:VARCHAR(12);not null;unique"`
	UserID       string `gorm:"type:VARCHAR(12)"`
	User         SQLiteUser
	CurrencyCode string              `gorm:"type:VARCHAR(3)"`
	Balance      float32             `gorm:"type:DECIMAL(10, 2)"`
	Transactions []SQLiteTransaction `gorm:"foreignKey:WalletID;references:ID"`
}

func (w *SQLiteWallet) CalculateTotalBalance() {
	var total float32
	for _, transaction := range w.Transactions {
		total += float32(transaction.Amount)
	}
	w.Balance = total
}

func NewSQLiteWalletRepository(db *gorm.DB) *SQLiteWalletRepository {
	db.AutoMigrate(&SQLiteWallet{})
	db.AutoMigrate(&SQLiteTransaction{})
	return &SQLiteWalletRepository{
		db: db,
	}
}

func (r *SQLiteWalletRepository) Create(createWalletDTO *models.CreateWalletDTO) (*models.Wallet, error) {
	wallet := &SQLiteWallet{
		ID:           generators.ID("wa"),
		CurrencyCode: createWalletDTO.CurrencyCode,
		UserID:       createWalletDTO.UserID,
		Balance:      0.0,
		Transactions: []SQLiteTransaction{},
	}
	result := r.db.Create(&wallet)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return r.mapSQLiteWalletToWallet(*wallet), nil
	}
}

func (r *SQLiteWalletRepository) Get(ID string) (*models.Wallet, error) {
	var wallet SQLiteWallet
	result := r.db.Model(&SQLiteWallet{ID: ID}).Find(&wallet)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return r.mapSQLiteWalletToWallet(wallet), nil
	}
}

func (r *SQLiteWalletRepository) List(listWalletsDTO *models.ListWalletsDTO) (*models.PaginatedWalletResultDTO, error) {
	switch {
	case listWalletsDTO.Limit > 100:
		listWalletsDTO.Limit = 100
	case listWalletsDTO.Limit <= 0:
		listWalletsDTO.Limit = 10
	}
	var wallets []SQLiteWallet
	var count int64

	result := db.Model(&SQLiteWallet{}).Offset((listWalletsDTO.Page - 1) * listWalletsDTO.Limit).Limit(listWalletsDTO.Limit).Find(&wallets)

	if result.Error != nil {
		return nil, result.Error
	}

	countResult := db.Model(&SQLiteWallet{}).Count(&count)

	if countResult.Error != nil {
		return nil, countResult.Error
	}

	data := r.mapSQLiteWalletsToWallets(&wallets)

	return &models.PaginatedWalletResultDTO{
		Page:       listWalletsDTO.Page,
		PerPage:    listWalletsDTO.Limit,
		Data:       *data,
		Count:      count,
		TotalPages: int(math.Ceil(float64(count) / float64(listWalletsDTO.Limit))),
	}, nil
}

func (r *SQLiteWalletRepository) AddTransaction(walletID string, addTransactionDTO models.AddTransactionDTO) (w *models.Wallet, err error) {
	var wallet SQLiteWallet
	result := r.db.Model(&SQLiteWallet{ID: walletID}).Find(&wallet)

	if result.Error != nil {
		return nil, result.Error
	} else {
		err = db.Model(&wallet).Association("Transactions").Append(&SQLiteTransaction{
			ID:              generators.ID("trx"),
			WalletID:        walletID,
			Amount:          addTransactionDTO.Amount,
			PreviousBalance: wallet.Balance,
			CurrencyCode:    addTransactionDTO.CurrencyCode,
		})
		if err != nil {
			return nil, err
		}
		wallet.CalculateTotalBalance()
		r.db.Save(&wallet)
		return r.mapSQLiteWalletToWallet(wallet), nil
	}
}

func (r *SQLiteWalletRepository) Delete(ID string) error {
	return fmt.Errorf("not implemented")
}

func (r *SQLiteWalletRepository) mapSQLiteWalletsToWallets(sqliteWallets *[]SQLiteWallet) *[]models.Wallet {
	data := &[]models.Wallet{}
	for _, wallet := range *sqliteWallets {
		*data = append(*data, *r.mapSQLiteWalletToWallet(wallet))
	}
	return data
}

func (r *SQLiteWalletRepository) mapSQLiteWalletToWallet(sqliteWallet SQLiteWallet) *models.Wallet {
	return &models.Wallet{
		ID:           sqliteWallet.ID,
		UserID:       sqliteWallet.UserID,
		CurrencyCode: sqliteWallet.CurrencyCode,
		Balance:      float64(sqliteWallet.Balance),
		Transactions: r.mapSQLiteTransactionsToTransactions(sqliteWallet.Transactions),
	}
}

func (r *SQLiteWalletRepository) mapSQLiteTransactionsToTransactions(sqliteTransactions []SQLiteTransaction) []models.Transaction {
	transactions := []models.Transaction{}
	for _, t := range sqliteTransactions {
		transactions = append(transactions, r.mapSQLiteTransactionToTransaction(t))
	}
	return transactions
}

func (r *SQLiteWalletRepository) mapSQLiteTransactionToTransaction(sqliteTransaction SQLiteTransaction) models.Transaction {
	return models.Transaction{
		ID:              sqliteTransaction.ID,
		WalletID:        sqliteTransaction.WalletID,
		Amount:          float64(sqliteTransaction.Amount),
		PreviousBalance: float64(sqliteTransaction.PreviousBalance),
		CurrencyCode:    sqliteTransaction.CurrencyCode,
	}
}
