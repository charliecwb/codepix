package repository

import (
	"errors"
	"github.com/charliecwb/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	res := r.Db.Create(transaction)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	res := r.Db.Save(transaction)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction *model.Transaction
	res := r.Db.Preload("AccountFrom.Bank").
		First(transaction, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	} else if transaction.ID == "" {
		return nil, errors.New("no transaction was found")
	}
	return transaction, nil
}
