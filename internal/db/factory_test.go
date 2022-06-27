package db

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/Keyhenge/PowerFlexChallenge/internal/model"
	"github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

var (
	badErr = errors.New("bad")
)

func Test_Factory_GetAll(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	db, dbMock, _ := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
	log := zap.NewNop().Sugar()

	expectedFactory := model.Factory{
		FactoryId: 1,
		ChartData: model.ChartData{
			SprocketProductionActual: []int32{1, 2, 3},
			SprocketProductionGoal:   []int32{2, 3, 4},
			Time:                     []int32{3, 4, 5},
		},
	}

	// TODO: sqlxmock currently doesn't have a way to mock selecting arrays AFAIK
	t.Run("Happy Path", func(t *testing.T) {
		factoryDB := FactoryDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		expectedFactories := []*model.Factory{&expectedFactory}
		rows := sqlxmock.NewRows([]string{"factory_id", "sprocket_production_actual", "sprocket_production_goal", "time"}).
			AddRow(1, pq.Array([]int32{1, 2, 3}), pq.Array([]int32{2, 3, 4}), pq.Array([]int32{3, 4, 5}))
		dbMock.ExpectQuery(getAllFactoriesQry).WillReturnRows(rows)

		result, err := factoryDB.GetAll(ctx)
		assert.Nil(err)
		assert.Equal(expectedFactories[0].FactoryId, result[0].FactoryId)
	})
	t.Run("No rows", func(t *testing.T) {
		factoryDB := FactoryDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectQuery(getAllFactoriesQry).WillReturnError(sql.ErrNoRows)

		result, err := factoryDB.GetAll(ctx)
		assert.Nil(err)
		assert.Nil(result)
	})
	t.Run("Error in query", func(t *testing.T) {
		factoryDB := FactoryDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectQuery(getAllFactoriesQry).WillReturnError(badErr)

		_, err := factoryDB.GetAll(ctx)
		assert.Equal(err, badErr)
	})
}

func Test_Factory_GetById(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	db, dbMock, _ := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
	log := zap.NewNop().Sugar()

	// TODO: sqlxmock currently doesn't have a way to mock selecting arrays AFAIK
	t.Run("Happy Path", func(t *testing.T) {
		factoryDB := FactoryDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		rows := sqlxmock.NewRows([]string{"sprocket_production_actual", "sprocket_production_goal", "time"}).
			AddRow(pq.Array([]int32{1, 2, 3}), pq.Array([]int32{2, 3, 4}), pq.Array([]int32{3, 4, 5}))
		dbMock.ExpectQuery(getFactoryByIdQry).WithArgs(1).WillReturnRows(rows)

		_, err := factoryDB.GetById(ctx, 1)
		assert.Nil(err)
		//assert.Equal(expectedFactory, *result)
	})
	t.Run("No rows", func(t *testing.T) {
		factoryDB := FactoryDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectQuery(getFactoryByIdQry).WithArgs(1).WillReturnError(sql.ErrNoRows)

		result, err := factoryDB.GetById(ctx, 1)
		assert.Nil(err)
		assert.Nil(result)
	})
	t.Run("Error in query", func(t *testing.T) {
		factoryDB := FactoryDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectQuery(getFactoryByIdQry).WithArgs(1).WillReturnError(badErr)

		_, err := factoryDB.GetById(ctx, 1)
		assert.Equal(badErr, err)
	})
}

func Test_Factory_New(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	db, dbMock, _ := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
	log := zap.NewNop().Sugar()

	factory := model.Factory{
		FactoryId: 1,
		ChartData: model.ChartData{
			SprocketProductionActual: []int32{1, 2, 3},
			SprocketProductionGoal:   []int32{2, 3, 4},
			Time:                     []int32{3, 4, 5},
		},
	}

	t.Run("Happy Path", func(t *testing.T) {
		factoryDB := FactoryDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		rows := sqlxmock.NewRows([]string{"factory_id"}).AddRow(factory.FactoryId)
		dbMock.ExpectQuery(createFactory).
			WithArgs(pq.Array(factory.ChartData.SprocketProductionActual),
				pq.Array(factory.ChartData.SprocketProductionGoal),
				pq.Array(factory.ChartData.Time)).
			WillReturnRows(rows)

		result, err := factoryDB.New(ctx, &factory)
		assert.Nil(err)
		assert.Equal(result, factory.FactoryId)
	})
	t.Run("Error in query", func(t *testing.T) {
		factoryDB := FactoryDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectQuery(createFactory).
			WithArgs(pq.Array(factory.ChartData.SprocketProductionActual),
				pq.Array(factory.ChartData.SprocketProductionGoal),
				pq.Array(factory.ChartData.Time)).
			WillReturnError(badErr)

		_, err := factoryDB.New(ctx, &factory)
		assert.Equal(badErr, err)
	})
}
