package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base      `valid:"required" json:"base"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `json:"bank" valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `json:"pix_keys" valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(ownername, number string, bank *Bank) (*Account, error) {
	account := &Account{
		OwnerName: ownername,
		Number:    number,
		Bank:      bank,
	}
	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return account, nil
}
