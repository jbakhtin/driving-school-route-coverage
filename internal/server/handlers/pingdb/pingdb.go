package pingdb

import "net/http"

type Handler struct {

}

func NewHandler() (Handler, error) {
	return Handler{

	}, nil
}

func (h *Handler) PingPostgres() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}

func (h *Handler) PingMongoDB() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}