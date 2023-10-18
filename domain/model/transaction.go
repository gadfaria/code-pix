package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionCanceled  string = "canceled"
	TransactionConfirmed string = "confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;not null" valid:"-"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyIDTo        string   `gorm:"column:pix_key_id_to;type:uuid;not null" valid:"-"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"-"`
	Status            string   `json:"status" gore:"type:varchar(20)" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)
	if err != nil {
		return err
	}

	if transaction.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if transaction.Status != TransactionPending && transaction.Status != TransactionCompleted && transaction.Status != TransactionConfirmed && transaction.Status != TransactionCanceled {
		return errors.New("invalid status for the transaction")
	}

	if transaction.PixKeyTo.AccountID == transaction.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Description: description,
		Status:      TransactionPending,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (transaction *Transaction) Complete() error {
	transaction.Status = TransactionCompleted
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	return err
}

func (transaction *Transaction) Confirm() error {
	transaction.Status = TransactionConfirmed
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()
	return err
}

func (transaction *Transaction) Cancel(description string) error {
	transaction.Status = TransactionCanceled
	transaction.UpdatedAt = time.Now()
	transaction.CancelDescription = description

	err := transaction.isValid()
	return err
}
