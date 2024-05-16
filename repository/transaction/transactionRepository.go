package transaction

import (
	"foodresq/entities"

	"gorm.io/gorm"
)

type TransactionRepositoryDB struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{
		db: db,
	}
}

func (repo *TransactionRepositoryDB) GetProductByName(productName string) (entities.RestaurantProduct, error) {
	var product entities.RestaurantProduct
	err := repo.db.Where("product_name = ?", productName).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repo *TransactionRepositoryDB) CreateTransaction(transaction entities.Transaction) (entities.Transaction, error) {
	err := repo.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *TransactionRepositoryDB) GetTransactionByUniqueCode(uniqueCode string) (entities.Transaction, error) {
	var transaction entities.Transaction
	if err := r.db.Where("unique_code = ?", uniqueCode).First(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *TransactionRepositoryDB) GetProductByID(productID uint) (entities.RestaurantProduct, error) {
	var product entities.RestaurantProduct
	if err := r.db.First(&product, productID).Error; err != nil {
		return product, err
	}
	return product, nil
}
