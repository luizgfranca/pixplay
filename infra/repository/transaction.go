package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/luizgfranca/pixplay/domain/model"
)

// type TransactionRepositoryInterface interface {
// 	Register(t *Transaction) error
// 	Save(t *Transaction) error
// 	Find(id string) (*Transaction, error)
// }

type TransactionRepositoryDB struct {
	Db *gorm.DB
}

func (r TransactionRepositoryDB) Register(t *model.Transaction) error {
	return r.Db.Create(t).Error
}

func (r TransactionRepositoryDB) Save(t *model.Transaction) error {
	return r.Db.Save(t).Error
}

func (r TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var t model.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&t, "id = ?", id)

	if t.ID == "" {
		return nil, fmt.Errorf("no register found tor this key")
	}

	return &t, nil
}
