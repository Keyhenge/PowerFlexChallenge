package api

import (
	"net/http"

	"github.com/Keyhenge/PowerFlexChallenge/internal/service"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type API struct {
	Factory  service.IFactoryService
	Sprocket service.ISprocketService
	log      zap.SugaredLogger
}

func (api *API) NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Route("/factory", func(r chi.Router) {
			r.Get("/all", api.FactoryGetAll())
			r.Get("/{factoryId}", api.FactoryGetById())
			r.Post("/new", api.FactoryNew())
		})
		r.Route("/sprocket", func(r chi.Router) {
			r.Get("/{sprocketId}", api.SprocketGetById())
			r.Post("/new", api.SprocketNew())
			r.Put("/update", api.SprocketUpdate())
		})
	})

	return router
}
