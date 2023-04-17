package points

import "net/http"

type Handler struct {

}

func NewHandler() (Handler, error) {
	return Handler{

	}, nil
}

func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}

func (h *Handler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}