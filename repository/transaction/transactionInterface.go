package transaction

import "foodresq/entities"

type TransactionInterface interface {
	GetProductByName(productName string) (entities.RestaurantProduct, error)
	CreateTransaction(transaction entities.Transaction) (entities.Transaction, error)
	GetTransactionByUniqueCode(uniqueCode string) (entities.Transaction, error)
	GetProductByID(productID uint) (entities.RestaurantProduct, error)
}
