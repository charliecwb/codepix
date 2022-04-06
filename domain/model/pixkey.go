package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) error
	FindKeyByKind(key, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
	FindBank(id string) (*Bank, error)
}

type PixKey struct {
	Base      `json:"base" valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `json:"account" valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixkey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixkey)
	if err != nil {
		return err
	}

	if pixkey.Kind != "email" && pixkey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pixkey.Status != "active" && pixkey.Status != "inactive" {
		return errors.New("invalid status")
	}

	return nil
}

func NewPixKey(kind, key string, account *Account) (*PixKey, error) {
	pixKey := &PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
