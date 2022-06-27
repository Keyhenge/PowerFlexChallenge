package model

import "net/http"

type Sprocket struct {
	SprocketId      int64 `json:"sprocket_id" db:"sprocket_id"`
	Teeth           int32 `json:"teeth" db:"teeth"`
	PitchDiameter   int32 `json:"pitch_diameter" db:"pitch_diameter"`
	OutsideDiameter int32 `json:"outside_diameter" db:"outside_diameter"`
	Pitch           int32 `json:"pitch" db:"pitch"`
}

func (s *Sprocket) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Sprocket) Bind(r *http.Request) error {
	return nil
}
