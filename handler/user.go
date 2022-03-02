package handler

import (
	"net/http"
	"refactoring/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user model.CreateUserRequest

	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	id, err := h.services.User.Create(user)
	if err != nil {
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})

}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.services.User.Get(id)

	if err != nil {
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, user)
}

func (h *Handler) searchUsers(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("s")

	list := h.services.User.Search(searchQuery)

	render.JSON(w, r, list)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	var userUpdate model.UpdateUserRequest

	err := render.DecodeJSON(r.Body, &userUpdate)
	if err != nil {
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	if err = h.services.User.Update(id, userUpdate); err != nil {
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.services.User.Delete(id); err != nil {
		_ = render.Render(w, r, model.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
