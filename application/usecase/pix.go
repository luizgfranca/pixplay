package usecase

import (
	"github.com/luizgfranca/pixplay/domain/model"
)

type PixKeyUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixKeyUseCase) RegisterKey(key string, kind string, accoundId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accoundId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	return p.PixKeyRepository.RegisterKey(pixKey)
}

func (p PixKeyUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	return p.PixKeyRepository.FindKeyByKind(key, kind)
}
