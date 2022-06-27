package model

import "net/http"

type Factory struct {
	FactoryId int64     `json:"factory_id" db:"factory_id"`
	ChartData ChartData `json:"chart_data" db:"chart_data"`
}

type ChartData struct {
	SprocketProductionActual []int32 `json:"sprocket_production_actual" db:"sprocket_production_actual"`
	SprocketProductionGoal   []int32 `json:"sprocket_production_goal" db:"sprocket_production_goal"`
	Time                     []int32 `json:"time" db:"time"`
}

func (s *Factory) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Factory) Bind(r *http.Request) error {
	return nil
}

type Factories struct {
	Factories []*Factory `json:"factories"`
}

func (s *Factories) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
