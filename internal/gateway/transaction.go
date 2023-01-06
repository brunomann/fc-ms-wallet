package gateway

import "github.com/brunomann/fc-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
