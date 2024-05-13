package transaction

import "foodresq/entities"

type TransactionInterface interface {
	GetProductByName(productName string) (entities.RestaurantProduct, error)
	InsertTransaction(transaction entities.Transaction) error
}
