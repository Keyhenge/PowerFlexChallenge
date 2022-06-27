package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Keyhenge/PowerFlexChallenge/internal/model"
	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func Test_Sprocket_GetById(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	db, dbMock, _ := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
	log := zap.NewNop().Sugar()

	expectedSprocket := model.Sprocket{
		Teeth:           1,
		PitchDiameter:   2,
		OutsideDiameter: 3,
		Pitch:           4,
	}

	// TODO: sqlxmock currently doesn't have a way to mock selecting arrays AFAIK
	t.Run("Happy Path", func(t *testing.T) {
		sprocketDB := SprocketDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		rows := sqlxmock.NewRows([]string{"teeth", "pitch_diameter", "outside_diameter", "pitch"}).AddRow(1, 2, 3, 4)
		dbMock.ExpectQuery(getSprocketByIdQry).WithArgs(1).WillReturnRows(rows)

		result, err := sprocketDB.GetById(ctx, 1)
		assert.Nil(err)
		assert.Equal(expectedSprocket, *result)
	})
	t.Run("No rows", func(t *testing.T) {
		sprocketDB := SprocketDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectQuery(getSprocketByIdQry).WithArgs(1).WillReturnError(sql.ErrNoRows)

		result, err := sprocketDB.GetById(ctx, 1)
		assert.Nil(err)
		assert.Nil(result)
	})
	t.Run("Error in query", func(t *testing.T) {
		sprocketDB := SprocketDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectQuery(getSprocketByIdQry).WithArgs(1).WillReturnError(badErr)

		_, err := sprocketDB.GetById(ctx, 1)
		assert.Equal(badErr, err)
	})
}

func Test_Sprocket_New(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	db, dbMock, _ := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
	log := zap.NewNop().Sugar()

	sprocket := model.Sprocket{
		SprocketId:      1,
		Teeth:           1,
		PitchDiameter:   2,
		OutsideDiameter: 3,
		Pitch:           4,
	}

	t.Run("Happy Path", func(t *testing.T) {
		sprocketDB := SprocketDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		rows := sqlxmock.NewRows([]string{"sprocket_id"}).AddRow(sprocket.SprocketId)
		dbMock.ExpectQuery(createSprocket).
			WithArgs(1, 2, 3, 4).
			WillReturnRows(rows)

		result, err := sprocketDB.New(ctx, &sprocket)
		assert.Nil(err)
		assert.Equal(result, sprocket.SprocketId)
	})
	t.Run("Error in query", func(t *testing.T) {
		sprocketDB := SprocketDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectQuery(createSprocket).
			WithArgs(1, 2, 3, 4).
			WillReturnError(badErr)

		_, err := sprocketDB.New(ctx, &sprocket)
		assert.Equal(badErr, err)
	})
}

func Test_Sprocket_Update(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	db, dbMock, _ := sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
	log := zap.NewNop().Sugar()

	sprocket := model.Sprocket{
		SprocketId:      1,
		Teeth:           1,
		PitchDiameter:   2,
		OutsideDiameter: 3,
		Pitch:           4,
	}

	t.Run("Happy Path", func(t *testing.T) {
		sprocketDB := SprocketDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectExec(updateSprocket).WithArgs(1, 2, 3, 4, 1).WillReturnResult(sqlxmock.NewResult(1, 1))

		err := sprocketDB.Update(ctx, &sprocket)
		assert.Nil(err)
	})
	t.Run("Error in query", func(t *testing.T) {
		sprocketDB := SprocketDB{
			BaseDB: BaseDB{
				DB:  db,
				Log: log,
			},
		}

		dbMock.ExpectExec(updateSprocket).WithArgs(1, 2, 3, 4, 1).WillReturnError(badErr)

		err := sprocketDB.Update(ctx, &sprocket)
		assert.Equal(badErr, err)
	})
}
