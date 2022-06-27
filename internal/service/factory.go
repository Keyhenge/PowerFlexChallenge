package service

import (
	"context"

	"github.com/Keyhenge/PowerFlexChallenge/internal/db"
	"github.com/Keyhenge/PowerFlexChallenge/internal/model"
)

type IFactoryService interface {
	GetAll(ctx context.Context) (*model.Factories, error)
	GetById(ctx context.Context, factoryId int64) (*model.Factory, error)
	New(ctx context.Context, factory *model.Factory) (int64, error)
}

type FactoryService struct {
	FactoryDB db.IFactoryDB
}

func (s *FactoryService) GetAll(ctx context.Context) (*model.Factories, error) {
	factories, err := s.FactoryDB.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Factories{Factories: factories}, nil

}

func (s *FactoryService) GetById(ctx context.Context, factoryId int64) (*model.Factory, error) {
	return s.FactoryDB.GetById(ctx, factoryId)
}

func (s *FactoryService) New(ctx context.Context, factory *model.Factory) (int64, error) {
	return s.FactoryDB.New(ctx, factory)
}
