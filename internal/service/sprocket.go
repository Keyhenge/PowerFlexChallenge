package service

import (
	"context"

	"github.com/Keyhenge/PowerFlexChallenge/internal/db"
	"github.com/Keyhenge/PowerFlexChallenge/internal/model"
)

type ISprocketService interface {
	GetById(ctx context.Context, sprocketId int64) (*model.Sprocket, error)
	New(ctx context.Context, sprocket *model.Sprocket) (int64, error)
	Update(ctx context.Context, sprocket *model.Sprocket) error
}

type SprocketService struct {
	SprocketDB db.ISprocketDB
}

func (s *SprocketService) GetById(ctx context.Context, sprocketId int64) (*model.Sprocket, error) {
	return s.SprocketDB.GetById(ctx, sprocketId)
}

func (s *SprocketService) New(ctx context.Context, sprocket *model.Sprocket) (int64, error) {
	return s.SprocketDB.New(ctx, sprocket)
}

func (s *SprocketService) Update(ctx context.Context, sprocket *model.Sprocket) error {
	return s.SprocketDB.Update(ctx, sprocket)
}
