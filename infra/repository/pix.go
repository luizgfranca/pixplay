package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/luizgfranca/pixplay/domain/model"
)

type PixKeyRepositoryDB struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDB) AddBank(bank *model.Bank) error {
	return r.Db.Create(bank).Error
}

func (r PixKeyRepositoryDB) AddAccount(account *model.Account) error {
	return r.Db.Create(account).Error
}

func (r PixKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return pixKey, nil
	}
	return nil, err
}

func (r PixKeyRepositoryDB) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no register found tor this key")
	}

	return &pixKey, nil
}

func (r PixKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no register found tor this key")
	}

	return &account, nil
}

func (r PixKeyRepositoryDB) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	r.Db.First(&bank, "id = ?", id)
	if bank.ID == "" {
		return nil, fmt.Errorf("no register found tor this key")
	}
	return &bank, nil
}
