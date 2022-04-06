package usecase

import (
	"errors"
	"github.com/charliecwb/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key, kind, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(key, kind, account)
	if err != nil {
		return nil, err
	}

	err = p.PixKeyRepository.RegisterKey(pixKey)
	if err != nil {
		return nil, err
	} else if pixKey.ID == "" {
		return nil, errors.New("no key created")
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(key, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
