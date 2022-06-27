package api

import (
	"net/http"
	"strconv"

	"github.com/Keyhenge/PowerFlexChallenge/internal/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

//Get sprocket in DB matching provided SprocketID
func (api *API) SprocketGetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		idString := chi.URLParam(r, "sprocketId")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			render.Render(w, r, ErrInternalServiceError(err))
			return
		}

		sprocket, err := api.Sprocket.GetById(ctx, id)
		if sprocket == nil {
			render.Render(w, r, ErrNotFound(err))
			return
		} else if err != nil {
			render.Render(w, r, ErrInternalServiceError(err))
			return
		}

		if err = render.Render(w, r, sprocket); err != nil {
			render.Render(w, r, ErrRender(err))
		}
		render.Status(r, 200)
	}
}

//Create Sprocket in DB, returning sprocket_id
func (api *API) SprocketNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := &model.Sprocket{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		sprocketId, err := api.Sprocket.New(ctx, data)
		if err != nil {
			render.Render(w, r, ErrInternalServiceError(err))
			return
		}

		if err = render.Render(w, r, &model.IdResponse{ID: sprocketId}); err != nil {
			render.Render(w, r, ErrRender(err))
		}
		render.Status(r, 200)
	}
}

//Update Sprocket in DB
func (api *API) SprocketUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := &model.Sprocket{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		err := api.Sprocket.Update(ctx, data)
		if err != nil {
			render.Render(w, r, ErrInternalServiceError(err))
			return
		}

		render.Status(r, 200)
	}
}
