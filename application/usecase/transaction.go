package usecase

import (
	"errors"
	"log"

	"github.com/luizgfranca/pixplay/domain/model"
)

type TransactionUseCase struct {
	PixKeyRepository      model.PixKeyRepositoryInterface
	TransactionRepository model.TransactionRepositoryInterface
}

func (t *TransactionUseCase) Register(
	accountId string,
	amount float64,
	pixKeyTo string,
	pixKeyKindTo string,
	description string,
) (*model.Transaction, error) {
	acc, err := t.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixKeyRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(acc, amount, pixKey)
	if err != nil {
		return nil, err
	}

	t.TransactionRepository.Save(transaction)
	if transaction.ID != "" {
		return transaction, nil
	}

	return nil, errors.New("Error persisting the transaction")
}

func (t *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	err = transaction.Confirm()
	if err != nil {
		return nil, err
	}

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	err = transaction.Complete()
	if err != nil {
		return nil, err
	}

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	err = transaction.Cancel(reason)
	if err != nil {
		return nil, err
	}

	err = t.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
