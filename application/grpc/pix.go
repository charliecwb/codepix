package grpc

import (
	"context"
	"github.com/charliecwb/codepix/application/grpc/pb"
	"github.com/charliecwb/codepix/application/usecase"
)

type PixGrpcService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

func NewPixGrpcService(useCase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{PixUseCase: useCase}
}

func (p *PixGrpcService) RegisterPixKey(ctx *context.Context, input *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.PixUseCase.RegisterKey(input.Key, input.Kind, input.AccountId)
	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Status: "created",
		Id:     key.ID,
	}, nil
}

func (p *PixGrpcService) Find(ctx *context.Context, input *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.PixUseCase.FindKey(input.Key, input.Kind)
	if err != nil {
		return nil, err
	}

	return &pb.PixKeyInfo{
		Id:   pixKey.ID,
		Kind: pixKey.Kind,
		Account: &pb.Account{
			AccountId:     pixKey.Account.ID,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankID,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		Key:       pixKey.Key,
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}
