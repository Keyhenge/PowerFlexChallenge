package api

import (
	"net/http"
	"strconv"

	"github.com/Keyhenge/PowerFlexChallenge/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

//Get all factories in DB
func (api *API) FactoryGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		factories, err := api.Factory.GetAll(ctx)
		if factories == nil {
			render.Render(w, r, ErrNotFound(err))
			return
		} else if err != nil {
			render.Render(w, r, ErrInternalServiceError(err))
			return
		}

		if err = render.Render(w, r, factories); err != nil {
			render.Render(w, r, ErrRender(err))
		}
		render.Status(r, 200)
	}
}

//Get factory in DB matching provided FactoryID
func (api *API) FactoryGetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		idString := chi.URLParam(r, "factoryId")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			render.Render(w, r, ErrInternalServiceError(err))
			return
		}

		factory, err := api.Factory.GetById(ctx, id)
		if factory == nil {
			render.Render(w, r, ErrNotFound(err))
			return
		} else if err != nil {
			render.Render(w, r, ErrInternalServiceError(err))
			return
		}

		if err = render.Render(w, r, factory); err != nil {
			render.Render(w, r, ErrRender(err))
		}
		render.Status(r, 200)
	}
}

//Create Factory in DB, returning factory_id
func (api *API) FactoryNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := &model.Factory{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		factoryId, err := api.Factory.New(ctx, data)
		if err != nil {
			render.Render(w, r, ErrInternalServiceError(err))
			return
		}

		if err = render.Render(w, r, &model.IdResponse{ID: factoryId}); err != nil {
			render.Render(w, r, ErrRender(err))
		}
		render.Status(r, 200)
	}
}
