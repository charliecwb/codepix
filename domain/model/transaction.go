package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	TransactionPending   = "Pending"
	TransactionCompleted = "Completed"
	TransactionError     = "Error"
	TransactionConfirmed = "Confirmed"
)

type TransactionRepositoryInterface interface {
	RegisterTransaction(transaction *Transaction) error
	Save(transactions *Transactions) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []*Transaction
}

type Transaction struct {
	Base              `json:"base" valid:"required"`
	AccountFrom       *Account `json:"accountFrom" valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `json:"pixKeyTo" valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancelDescription" valid:"-"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)
	if err != nil {
		return err
	}

	if t.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if t.Status != TransactionCompleted && t.Status != TransactionConfirmed && t.Status != TransactionError &&
		t.Status != TransactionPending {
		return errors.New("invalid status for the transaction")
	}

	if t.PixKeyTo.Account.ID == t.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := &Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.CancelDescription = description
	err := t.isValid()
	return err
}
