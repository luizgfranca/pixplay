package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "PENDING"
	TransactionCompleted string = "COMPLETED"
	TransactionError     string = "ERROR"
	TransactionConfirmed string = "CONFIRMED"
)

type TransactionRepositoryInterface interface {
	Register(t *Transaction) error
	Save(t *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transaction struct {
	Base              `valid:"required"`
	AccountIdFrom     string   `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyIdTo        string   `gorm:"column:pix_key_id_to;type:uuid;" valid:"-"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"descri ption" gorm:"type:varchar(255)" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

type Transactions struct {
	Transaction []Transaction
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("transaction value is less than zero")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionConfirmed && t.Status != TransactionError {
		return errors.New("invalid transaction status")
	}

	if t.AccountFrom == t.PixKeyTo.Account {
		return errors.New("destination can't be the same account")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey) (*Transaction, error) {

	t := Transaction{
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
	}

	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()

	err := t.isValid()

	return &t, err
}

func (t *Transaction) updateStatus(newStatus string) error {
	t.Status = newStatus
	t.UpdatedAt = time.Now()
	return t.isValid()
}

func (t *Transaction) Complete() error {
	return t.updateStatus(TransactionCompleted)
}

func (t *Transaction) Confirm() error {
	return t.updateStatus(TransactionConfirmed)
}

func (t *Transaction) Cancel(description string) error {
	t.CancelDescription = description
	return t.updateStatus(TransactionError)
}
