package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Keyhenge/PowerFlexChallenge/internal/model"
	"github.com/lib/pq"
)

const (
	getAllFactoriesQry = `SELECT * FROM factory`
	getFactoryByIdQry  = `SELECT sprocket_production_actual
                                   , sprocket_production_goal
                                   , time 
                                  FROM factory WHERE factory_id = $1`
	createFactory = `INSERT INTO factory(sprocket_production_actual, sprocket_production_goal, time)
                             VALUES($1, $2, $3)
                             RETURNING factory_id;`
)

type IFactoryDB interface {
	GetAll(ctx context.Context) ([]*model.Factory, error)
	GetById(ctx context.Context, factoryId int64) (*model.Factory, error)
	New(ctx context.Context, factory *model.Factory) (int64, error)
}

type FactoryDB struct {
	BaseDB
}

func (db *FactoryDB) GetAll(ctx context.Context) ([]*model.Factory, error) {
	factories := []*model.Factory{}
	rows, err := db.DB.QueryContext(ctx, getAllFactoriesQry)
	if errors.Is(err, sql.ErrNoRows) {
		db.Log.Info("No factories in DB")
		return nil, nil
	} else if err != nil {
		db.Log.Errorw("FactoryDB - GetAll - Error while retrieving factories", "error", err)
		return nil, err
	}

	for rows.Next() {
		factory := model.Factory{}
		err = rows.Scan(
			&factory.FactoryId,
			pq.Array(&factory.ChartData.SprocketProductionActual),
			pq.Array(&factory.ChartData.SprocketProductionGoal),
			pq.Array(&factory.ChartData.Time),
		)
		if err != nil {
			db.Log.Errorw("FactoryDB - GetAll - Error while scanning factories", "error", err)
			return nil, err
		}

		factories = append(factories, &factory)
	}

	return factories, nil
}

func (db *FactoryDB) GetById(ctx context.Context, factoryId int64) (*model.Factory, error) {
	factory := model.Factory{}
	err := db.DB.QueryRowContext(ctx, getFactoryByIdQry, factoryId).Scan(
		pq.Array(&factory.ChartData.SprocketProductionActual),
		pq.Array(&factory.ChartData.SprocketProductionGoal),
		pq.Array(&factory.ChartData.Time),
	)
	if errors.Is(err, sql.ErrNoRows) {
		db.Log.Infow("Could not find factory for specified FactoryID", "factory_id", factoryId)
		return nil, nil
	} else if err != nil {
		db.Log.Errorw("FactoryDB - GetById - Error while retrieving factory", "error", err, "factory_id", factoryId)
		return nil, err
	}

	return &factory, nil
}

func (db *FactoryDB) New(ctx context.Context, factory *model.Factory) (int64, error) {
	var factoryId int64
	err := db.DB.GetContext(ctx, &factoryId, createFactory,
		pq.Array(factory.ChartData.SprocketProductionActual),
		pq.Array(factory.ChartData.SprocketProductionGoal),
		pq.Array(factory.ChartData.Time),
	)
	if err != nil {
		db.Log.Errorw("FactoryDB - New - Error while creating factory", "error", err)
		return 0, err
	}

	return factoryId, nil
}
