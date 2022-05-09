package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string `json:"owner_name" valid:"notnull"`
	Bank      *Bank  `valid:"-"`
	Number    string `json:"number" valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}

func newAccount(bank *Bank, number string, ownerName string) (*Account, error) {
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
