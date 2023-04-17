package auth

import (
	"net/http"
)

type Handler struct {

}

func NewHandler() (Handler, error) {
	return Handler{

	}, nil
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}

func (h *Handler) LogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}

func (h *Handler) LogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}