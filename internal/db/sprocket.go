package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Keyhenge/PowerFlexChallenge/internal/model"
)

const (
	getSprocketByIdQry = `SELECT * FROM sprocket WHERE sprocket_id = $1;`
	createSprocket     = `INSERT INTO sprocket(teeth, pitch_diameter, outside_diameter, pitch)
                                  VALUES($1, $2, $3, $4)
                                  RETURNING sprocket_id;`
	updateSprocket = `UPDATE sprocket
                              SET teeth = $1,
                                  pitch_diameter = $2,
                                  outside_diameter = $3,
                                  pitch = $4
                              WHERE sprocket_id = $5`
)

type ISprocketDB interface {
	GetById(ctx context.Context, sprocketId int64) (*model.Sprocket, error)
	New(ctx context.Context, sprocket *model.Sprocket) (int64, error)
	Update(ctx context.Context, sprocket *model.Sprocket) error
}

type SprocketDB struct {
	BaseDB
}

func (db *SprocketDB) GetById(ctx context.Context, sprocketId int64) (*model.Sprocket, error) {
	sprocket := &model.Sprocket{}
	err := db.DB.GetContext(ctx, sprocket, getSprocketByIdQry, sprocketId)
	if errors.Is(err, sql.ErrNoRows) {
		db.Log.Infow("Could not find sprocket for specified SprocketID", "sprocket_id", sprocketId)
		return nil, nil
	} else if err != nil {
		db.Log.Errorw("SprocketDB - GetById - Error while retrieving sprocket", "error", err, "sprocket_id", sprocketId)
		return nil, err
	}

	return sprocket, nil
}

func (db *SprocketDB) New(ctx context.Context, sprocket *model.Sprocket) (int64, error) {
	var sprocketId int64
	err := db.DB.GetContext(ctx, &sprocketId, createSprocket,
		sprocket.Teeth,
		sprocket.PitchDiameter,
		sprocket.OutsideDiameter,
		sprocket.Pitch,
	)
	if err != nil {
		db.Log.Errorw("SprocketDB - New - Error while creating sprocket", "error", err)
		return 0, err
	}

	return sprocketId, nil
}

func (db *SprocketDB) Update(ctx context.Context, sprocket *model.Sprocket) error {
	_, err := db.DB.ExecContext(ctx, updateSprocket,
		sprocket.Teeth,
		sprocket.PitchDiameter,
		sprocket.OutsideDiameter,
		sprocket.Pitch,
		sprocket.SprocketId,
	)
	if err != nil {
		db.Log.Errorw("SprocketDB - Update - Error while updating sprocket", "error", err)
	}

	return err
}
