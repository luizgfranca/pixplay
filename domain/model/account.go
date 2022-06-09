package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string `gorm:"column:owner_name;type:varchar(255);not null" json:"owner_name" valid:"notnull"`
	BankID    string `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Bank      *Bank  `valid:"-"`
	Number    string `json:"number" gorm:"type:varchar(20)" valid:"notnull"`

	PixKeys []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	acc := Account{
		Bank:      bank,
		Number:    number,
		OwnerName: ownerName,
	}

	acc.ID = uuid.NewV4().String()
	acc.CreatedAt = time.Now()
	acc.UpdatedAt = time.Now()

	err := acc.isValid()

	return &acc, err
}
