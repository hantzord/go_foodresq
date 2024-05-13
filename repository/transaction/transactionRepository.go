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

func (repository *TransactionRepositoryDB) GetProductByName(productName string) (entities.RestaurantProduct, error) {
	var product entities.RestaurantProduct
	if err := repository.db.Where("product_name = ?", productName).First(&product).Error; err != nil {
		return entities.RestaurantProduct{}, err
	}
	return product, nil
}

func (repository *TransactionRepositoryDB) InsertTransaction(transaction entities.Transaction) error {
	if err := repository.db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}
