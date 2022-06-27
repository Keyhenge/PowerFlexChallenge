package model

import "net/http"

type IdResponse struct {
	ID int64 `json:"id"`
}

func (s *IdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
