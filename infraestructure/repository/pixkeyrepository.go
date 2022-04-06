package repository

import (
	"errors"
	"github.com/charliecwb/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r *PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	res := r.Db.Create(bank)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) error {
	err := r.Db.Create(pixKey)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *PixKeyRepositoryDb) FindKeyByKind(key, kind string) (*model.PixKey, error) {
	var pixKey *model.PixKey
	res := r.Db.Preload("Account.Bank").
		Find(&model.PixKey{
			Key:  key,
			Kind: kind}).
		First(pixKey)
	if res.Error != nil {
		return nil, res.Error
	} else if pixKey.ID == "" {
		return nil, errors.New("no key was found")
	}
	return pixKey, nil
}

func (r *PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	res := r.Db.Create(account)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account *model.Account
	res := r.Db.Preload("Bank").
		First(account, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	} else if account.ID == "" {
		return nil, errors.New("no account was found")
	}
	return account, nil
}

func (r *PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank *model.Bank
	res := r.Db.First(bank, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	} else if bank.ID == "" {
		return nil, errors.New("no bank was found")
	}
	return bank, nil
}
